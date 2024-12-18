pipeline {
    agent { label 'centos' }

    environment {
        SONARQUBE_SCANNER_HOME = tool 'SonarScanner'
        OUTPUT_DIR = 'build'
        BRANCH_NAME = "${env.BRANCH}"
        GITHUB_REPO = "xiao-monitor/monitor-dashboard"
        GITHUB = credentials('github-token')
        TAG_VERSION = ''
    }

    stages {
        stage('基于分支执行操作') {
            steps {
                script {
                    if (BRANCH_NAME == 'null' || BRANCH_NAME == '') {
                        BRANCH_NAME = sh(script: "git rev-parse --abbrev-ref HEAD", returnStdout: true).trim()
                    } else if (BRANCH_NAME.startsWith('refs/tags/')) {
                        echo("❌ 不能在 Tag 分支上执行操作！")
                        return;
                    }
                }
                echo "当前分支是 ${BRANCH_NAME}，执行操作..."
            }
        }

        stage('签出代码') {
            when {
                expression { BRANCH_NAME == 'master' && !BRANCH_NAME.startsWith('refs/tags/') }
            }
            steps {
                echo "拉取代码中..."
                checkout scm
            }
        }

        stage('代码分析') {
            when {
                expression { !BRANCH_NAME.startsWith('refs/tags/') }
            }
            steps {
                echo "执行 SonarQube 分析..."
                withSonarQubeEnv('SonarScanner') {
                    sh "${SONARQUBE_SCANNER_HOME}/bin/sonar-scanner"
                }
            }
        }

        stage('项目打包') {
            when {
                expression { BRANCH_NAME == 'master' && !BRANCH_NAME.startsWith('refs/tags/') }
            }
            steps {
                script {
                    TAG_VERSION = sh(script: "cat ${WORKSPACE}/hack/config.yaml | grep 'version' | awk '{print \$2}'", returnStdout: true).trim()
                }
                echo "清空构建目录..."
                sh "rm -rf ${WORKSPACE}/build/${TAG_VERSION}"
                echo "编译项目..."
                sh '''
                    if ! command -v gf >/dev/null 2>&1; then
                        echo "gf 命令不存在，开始下载安装..."
                        wget -O gf "https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH)" && chmod +x gf && ./gf install -y && rm ./gf
                    else
                        echo "gf 命令已存在，跳过安装。"
                    fi
                '''
                sh 'gf build'
            }
        }

        stage('收集构建产物') {
            when {
                expression { BRANCH_NAME == 'master' && !BRANCH_NAME.startsWith('refs/tags/') }
            }
            steps {
                echo("获取版本号: ${TAG_VERSION}")
                echo "收集所有构建产物..."
                script {
                    def files = sh(script: "find ${OUTPUT_DIR} -type f -name 'monitor-dashboard*'", returnStdout: true).trim().split('\n')

                    files.each { file ->
                        def dirName = sh(script: "dirname ${file}", returnStdout: true).trim()
                        def newName = "${dirName}_${TAG_VERSION}_\$(basename ${file})"
                        echo "重命名文件: $file -> ${newName}"
                        sh("mv $file $newName")
                        sh("rm -rf ${dirName}")
                    }
                }
            }
        }

        stage('发布到 GitHub Release') {
            when {
                expression { BRANCH_NAME == 'master' && !BRANCH_NAME.startsWith('refs/tags/') }
            }
            steps {
                script {
                    def RELEASE_NAME = "RELEASE-${TAG_VERSION}"
                    def CHANGELOG_FILE = "CHANGELOG.md"

                    echo("生成 CHANGELOG...")
                    sh """
                        PREV_TAG=\$(git describe --tags --abbrev=0 HEAD^)
                        echo "# DESCRIPTION" > ${CHANGELOG_FILE}
                        echo "\$(git log -1 --pretty=format:'%B') \n" >> ${CHANGELOG_FILE}
                        echo "# CHANGELOG" >> ${CHANGELOG_FILE}
                        git log --pretty=format:'- %s (@%an) [%h]' \$PREV_TAG..HEAD >> ${CHANGELOG_FILE}
                        echo "" >> ${CHANGELOG_FILE}
                    """

                    def loginStatus = sh(
                        script: "echo $GITHUB | gh auth login --with-token && gh auth status | grep 'Logged in'",
                        returnStatus: true
                    )

                    if (loginStatus != 0) {
                        error "❌ GitHub CLI 登录失败，请检查 GITHUB_TOKEN 是否正确！"
                    } else {
                        echo "✅ GitHub CLI 登录成功"
                    }

                    sh """
                        echo "发布到 GitHub Release..."
                        gh release create ${TAG_VERSION} ${OUTPUT_DIR}/${TAG_VERSION}/** \\
                            --repo ${GITHUB_REPO} \\
                            --title "${RELEASE_NAME}" \\
                            --notes-file ${CHANGELOG_FILE}
                    """
                }
            }
        }
    }

    post {
        success {
            echo "流水线执行成功 🎉"
        }
        failure {
            echo "流水线执行失败，请检查日志 ❌"
        }
    }
}