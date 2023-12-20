pipeline{
    agent any
    tools {
        go 'go-1.21.5'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage("Test"){
            steps{
                git "https://github.com/hanshal101/jwt"
                sh 'go test ./...'
            }
        }
    }
    }