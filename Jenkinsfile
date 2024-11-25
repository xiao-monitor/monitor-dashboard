node {
  stage('SCM') {
    checkout scm
  }
  stage('SonarQube分析') {
    def scannerHome = tool 'SonarScanner';
    withSonarQubeEnv() {
      sh "${scannerHome}/bin/sonar-scanner"
    }
  }
}