pipeline {
    agent {label: "centos"}

    tools {
        sonarScanner 'SonarScanner'
    }

    stages {
        stage('SCM') {
            steps {
                echo "拉取代码..."
                checkout scm
            }
        }

        stage('SonarQube Analysis') {
            steps {
                withSonarQubeEnv('SonarScanner') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
    }

    post {
        success {
            echo(message: "SonarQube 分析完成 ✅")
        }
        failure {
            error(message: "SonarQube 分析失败，请检查日志 ❌")
        }
    }
}