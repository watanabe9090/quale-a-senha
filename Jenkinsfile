pipeline {
    agent any

    environment {
        GO_SERV_PORT = 'on'
    }
    stages {
        stage('test') {
            agent {
                docker {
                    image 'golang:1.23.1-alpine3.20' 
                }
            }
            steps {
                sh 'go test ./...'
            }
        }
        stage('build') {
            steps {
                sh 'go version'
            }
        }
        stage('pull') {
            steps {
                sh 'docker image push 192.168.18.112:5000/quale-a-senha'
            }
        }
    }
}