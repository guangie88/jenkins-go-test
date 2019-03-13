pipeline {
  agent {
    docker { image 'golang:1.12-alpine' }
  }
  stages {
    stage('Unit Test') {
      steps {
        sh 'go test -v ./...'
      }
    }
  }
}
