pipeline {
  agent {
    docker {
      image 'go:latest'
      args 'docker build -t hanshal785/jwt'
    }

  }
  stages {
    stage('Test') {
      steps {
        git(url: 'https://github.com/hanshal101/jwt', branch: 'main')
        sh 'go test ./...'
      }
    }

    stage('Go Build') {
      steps {
        sh 'go build .'
      }
    }

    stage('Docker Image Build') {
      steps {
        script {
          app = docker.build("hanshal785/jwt")
        }

      }
    }

    stage('Message') {
      steps {
        sh 'echo "Successfully ran the Jenkinsfile"'
      }
    }

  }
  tools {
    go 'go-1.21.5'
    git 'Default'
  }
  environment {
    GO111MODULE = 'on'
  }
}