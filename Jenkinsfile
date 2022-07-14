pipeline{
  agent any
  environment {
    uRLWEBHOOK = "https:/webhookb2/18b18bbb-c20e-4c25-9a90-71936337e05f@f1251ac1-12f6-435a-9795-54df89862df1/JenkinsCI/239a29e64c1b4ffe8fa324a1316648cb/a72bb36f-1a21-4c80-a746-5fe2a5005ee4"
    BRANCH_NAME = "${GIT_BRANCH.split("/")[1]}"
    REPOURL = "adapt-registry:5003/repository/goadapt/"
    registrycred = 'registrycred'
    dockerImage = ''
    BASE_DIR = "${pwd}"
    GO114MODULE = 'on'
    CGO_ENABLED = 0
    GOPATH1 = "/var/lib/jenkins/sample-ap/sample-app/cart/src"
    GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    EMAIL="shubham.tamboli1@gmail.com"
    TAG= "v1.1.1"
  }
  stages{
   stage("Workspace_cleanup"){
        //Cleaning WorkSpace
     steps{
        step([$class: 'WsCleanup'])
     }
   }	  
   stage("Repo_clone"){
       //Clone repo from GitHub
     steps {
          checkout ([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsId: 'shubham', url: 'git@github.com:shubh9975/sample-app.git']]])
	  }
   }	  
   stage("Hello Calsoft"){
	   steps {
		sh 'echo "Hello Everyone"'
	   }
     }
   stage('Pre dependaencies.') {
     steps {
        withEnv(["PATH+GO=${GOPATH}/bin"]){
           sh '''
	      export GO111MODULE="on"
              sudo docker system prune -f
              echo 'Installing dependencies'
              go version
	      go mod init cart/src
              go get -u golang.org/x/lint/golint
              go install github.com/securego/gosec/v2/cmd/gosec@latest
              go get -u github.com/securego/gosec/cmd/gosec
              export PATH="/usr/local/sonar-scanner/bin:$PATH"
           '''
        }
     }
   }
   stage("static analysis"){
     steps{
        withEnv(["PATH+GO=${GOPATH}/bin"]){
          //dir('cart/src'){
           sh '''
	     echo "Performing the Testcase on our code"
	     cd cart/src
	     go test -v
	     echo "Testing code"
             go run cart/src/main.go
           '''
           //}
        }
     }
   }
   stage("Image Building"){
     steps{
        script{
	  dir('cart/'){
          sh '''
	      sudo docker build -t myimage .
	  '''
	  }
        }
     }
   }
	  
  } 
}
  
