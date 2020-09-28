<h1 align="center"> Healer - Microservices Local Development </h1>
<p align="center">
 <img src="https://img.techpowerup.org/200925/0-1.png" width="400px">
</p>
<h2>manage microservices in local environment</h2>

:exclamation:  Currently, Healer is in the Pre-alpha version. :exclamation:

### Installation:

    go get https://github.com/ghaini/healer 
    
### commands:

**create a project**:
   
    command: healer project-create -n {project name}
    
    example: healer project-create -n myProject
    
**add a command to project**:
   
    command: healer command-create -p {project name} -c {command}
    
    example: healer command-create -p myProject -c "ls -al"
    
**up project**:
       
    command: healer project-up -p {project name}
        
    example: healer project-up -p myProject
        
    
**list of projects**:
       
    command: healer project-list
