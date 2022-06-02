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



## Accesing services outside the cluste

### Using minikube service

Returns a URL to connect to a service

```minikube -n NAMESPACE service SERVICE_NAME --url```

Example:

```minikube service -n dev multiple-operation-svc-go-kit-svc --url```

### Using Ingress (WIP)

#### Enable an Ingress Controller

If you want to access your services from outside the cluster, you will nedd and Ingress controller. You can enable the NGINX Ingress controller in Minikube running the following command:

```minikube addons enable ingress```

More information:

[Enable the Ingress controller](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)


#### Create default ingress class