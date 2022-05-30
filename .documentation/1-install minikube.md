# Install MiniKube

## Previous Requirements
To have Docker installed. In my case I used Docker Desktop for Windows
https://docs.docker.com/desktop/


# Installation

An easy way to install minikube is following de tutorial in the oficcial website.  
https://minikube.sigs.k8s.io/docs/start/


## Test installation

Run the comand 

```minikube start```

After finished you can check accesing the cluster.

```kubectl get po -A``` 

if you don have kubectl, you can use this command. 

``` minikube kubectl -- get po -A ```	

you can see how to install it in this page [Installing Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
