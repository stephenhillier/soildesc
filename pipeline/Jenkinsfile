pipeline {
  agent {
    docker {
      image 'golang:1.10-alpine'
    }
  }

  stages {
    stage('Build') {
      steps {
        sh 'go version'
      }
    }
  }
}
