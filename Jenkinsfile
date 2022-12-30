node {

    tools {
        go 'go'
    }


  stage('SCM') {
    checkout scm
  }

  stage('SonarQube Analysis') {
    def scannerHome = tool 'SonarScanner';
    withSonarQubeEnv() {
      sh "${scannerHome}/bin/sonar-scanner"
    }
  }

  stage("build") {
      echo 'BUILD EXECUTION STARTED'
      sh 'go version'
  }
}