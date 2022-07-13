pipeline{
  agent any
  stages{
   stage("Repo_clone"){
       //Clone repo from GitHub
     steps {
          checkout ([$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsId: 'shubham', url: 'git@github.com:shubh9975/sample-app.git']]])
	  }
   }	  
   stage("Hello Calsoft"){
     step echo "Hello Everyone"
     }
  } 
}
