#!groovy

node {

    stage('Checkout Code') {
        checkout scm: [
            $class: 'GitSCM',
            branches: [
                [name: "${env.BRANCH_NAME}"]
            ],
            userRemoteConfigs: [[credentialsId:'jenkins-github-ssh-pat', url: 'https://github.com/ernestngugi/sil-devops.git']]
        ]
    }

    stage("remove docker container") {
        dir ('api') {
            sh '''
            sudo docker-compose down --remove-orphans
        '''
        }
    }

    stage("build docker container") {
        dir ('api') {
            sh '''
            sudo docker-compose up -d
        '''
        }
    }

    // stage("remove sil-api image") {
    //     dir ('api') {
    //         sh '''
    //         sudo docker rm go-blog
    //     '''
    //     }
    // }

    stage("build sil-api image") {
        dir ('api') {
            sh '''
            sudo docker build -t go-blog .
        '''
        }
    }

    stage("run sil-api image") {
        dir ('api') {
            sh '''
            sudo docker run go-blog
        '''
        }
    }
}