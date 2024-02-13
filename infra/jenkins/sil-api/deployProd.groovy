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

    stage("build and deploy") {
        dir ('api') {
            sh '''
            sudo docker-compose up -d
        '''
        }
    }
}