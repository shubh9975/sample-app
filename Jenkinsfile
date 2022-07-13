pipeline{
  agent any
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
  } 
}
