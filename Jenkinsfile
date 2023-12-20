pipeline {
    agent any
    tools {
        go 'go-1.21.5'
        git 'Default'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage("Test") {
            steps {
                git url: 'https://github.com/hanshal101/jwt', branch: 'main'
                sh 'go test ./...'
            }
        }
        stage("Build") {
            steps {
                sh 'go build .'
            }
        }
    }
}