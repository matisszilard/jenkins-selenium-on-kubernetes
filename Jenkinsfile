#!/usr/bin/env groovy

pipeline {
    agent none
    stages {
        stage('Run tests in chrome') {
          agent {
            kubernetes {
              label 'selenium-chrome'
              defaultContainer 'builder'
              yamlFile 'test-env.yaml'
            } // kubernetes
          } // agent
          steps {
              script {
                sh "./wait-for-grid.sh go build -o run_test_cases && ./run_test_cases"
              } // container
          } // steps
        } // stage(build)
    } // stages
} // pipeline
