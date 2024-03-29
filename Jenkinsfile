pipeline {
  agent {
    label 'ubuntu_docker_label'
  }
  stages {
    stage("Lint") {
      steps {
        sh "make fmt && git diff --exit-code"
      }
    }
    stage("Test") {
      steps {
        sh "make test"
      }
    }
    stage("Build") {
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          sh "make docker push"
        }
      }
    }
    stage("Push") {
      when {
        branch "master"
      }
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          sh "make push IMAGE_VERSION=latest"
        }
      }
    }
    
    stage('Push charts') {
        steps {
          withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
            withAWS(region:'<insert-the-region-id>', credentials:'<insert-the-creds-id>') {
                sh "IMAGE_VERSION=\$(IMAGE_VERSION)-j$BUILD_NUMBER make push-chart"
              }
          }
          archiveArtifacts artifacts: 'helm/$(CHART_NAME)-*.tgz'
          archiveArtifacts artifacts: 'helm.properties'
        }
    }
  }
  post {
    cleanup {
      sh "make clean || true"
      cleanWs()
    }
  }
}
