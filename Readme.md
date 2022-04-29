I created this tutorial to understand better about CI/CD and Microservices Architecture.

The general idea is to deploy some microservices to a Kubernetes cluster, and start implementing 
Microservices Patterns, to know how use them in diferent situations.

# Tools Utilized

- Minikube. To have a kubernetes cluster.
- Github Actions (CI). To build images and put in docker hub
- ArgoCD and Helm (CD). To manake resources configuration and deploy  in the kubernetes cluster
- Go Kit. To develop the microservices


# Summary

[1 - Install Minikube](.documentation/1-install%20minikube.md)  
[2 - Install ArgoCD](.documentation/2-install%20argocd.md)  
[2.1 - Configure ArgoCD](.documentation/2.1-configure%20argocd.md)  
[3 - Github Action(CI)](.documentation/3-github%20action%20(CI).md)  
[4 - Configuration Manager with HELM](.documentation/4-configuration%20manager%20-%20helm.md)