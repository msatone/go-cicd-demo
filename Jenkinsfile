pipeline {
    agent any

    environment {
        GO111MODULE = "on"
        GOPROXY = "https://proxy.golang.org"

        AWS_REGION = "us-east-1"
        AWS_ACCOUNT_ID = "494934331459"

        IMAGE_NAME = "dummy"
        IMAGE_TAG = "latest"

        ECR_REGISTRY = "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
        FULL_IMAGE = "${ECR_REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
    }

    options {
        timestamps()
        ansiColor('xterm')
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    stages {

        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Version Check') {
            steps {
                sh '''
                    echo "Go Version"
                    go version
                '''
            }
        }

        stage('Install Dependencies') {
            steps {
                sh '''
                    chmod +x scripts/*.sh

                    go mod tidy

                    go mod download

                    if ! command -v golangci-lint >/dev/null 2>&1
                    then
                        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
                        | sh -s -- -b /usr/local/bin v1.55.2
                    fi
                '''
            }
        }

        stage('Quality Checks') {

            parallel {

                stage('Formatting') {
                    steps {
                        sh '''
                            echo "Formatting"
                            go fmt ./...
                        '''
                    }
                }

                stage('Lint') {
                    steps {
                        sh '''
                            chmod +x scripts/lint.sh
                            ./scripts/lint.sh
                        '''
                    }
                }

                stage('Unit Tests') {
                    steps {
                        sh '''
                            chmod +x scripts/test.sh
                            ./scripts/test.sh
                        '''
                    }
                }

            }
        }

        stage('Build Binary') {
            steps {

                sh '''
                    chmod +x scripts/build.sh
                    ./scripts/build.sh

                    ls -lah bin
                '''

            }
        }

        stage('Docker Build') {
            steps {

                sh '''
                    docker build \
                    -t ${IMAGE_NAME}:${IMAGE_TAG} \
                    -f deployment/Dockerfile .
                '''

            }
        }

        stage('Trivy Scan') {
            steps {

                sh '''
                    docker run --rm \
                    -v /var/run/docker.sock:/var/run/docker.sock \
                    aquasec/trivy image \
                    --severity HIGH,CRITICAL \
                    ${IMAGE_NAME}:${IMAGE_TAG}
                '''

            }
        }

        stage('Login to AWS ECR') {

            steps {

                withCredentials([
                    [$class: 'AmazonWebServicesCredentialsBinding',
                    credentialsId: 'aws-creds']
                ]) {

                    sh '''
                        aws ecr get-login-password \
                        --region ${AWS_REGION} \
                        | docker login \
                        --username AWS \
                        --password-stdin ${ECR_REGISTRY}
                    '''
                }

            }

        }

        stage('Push Image') {

            steps {

                sh '''

                    docker tag \
                    ${IMAGE_NAME}:${IMAGE_TAG} \
                    ${FULL_IMAGE}

                    docker push \
                    ${FULL_IMAGE}

                '''

            }

        }

        stage('Deploy to Kubernetes') {

            steps {

                withCredentials([
                    file(credentialsId: 'kubeconfig',
                    variable: 'KUBECONFIG')
                ]) {

                    sh '''

                        kubectl apply -f k8s/namespace.yaml

                        kubectl apply -f k8s/deployment.yaml

                        kubectl apply -f k8s/service.yaml

                        kubectl rollout status deployment/my-go-app -n my-app-ns

                    '''

                }

            }

        }

    }

    post {

        always {

            archiveArtifacts artifacts: 'bin/**', fingerprint: true

            archiveArtifacts artifacts: 'coverage.out', allowEmptyArchive: true

        }

        success {

            echo '================================='
            echo 'Build Successful'
            echo '================================='

        }

        failure {

            echo '================================='
            echo 'Build Failed'
            echo '================================='

        }
    }

}
