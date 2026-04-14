pipeline {
  agent any

  environment {
    IMAGE = "url-shortener:dev"
    NAMESPACE = "url-shortener"
  }

  stages {

    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build Docker image') {
      steps {
        sh 'docker build -t $IMAGE ./app'
      }
    }

    stage('Load image into kind') {
      steps {
        sh 'kind load docker-image $IMAGE --name devops-lab'
      }
    }

    stage('Deploy with Helm') {
      steps {
        sh 'helm upgrade --install url-shortener ./deploy/helm/url-shortener -n $NAMESPACE'
      }
    }
  }
}