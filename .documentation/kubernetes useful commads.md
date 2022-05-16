# Deploy a resource 
```kubectl apply -f .\app-project-calculator-dev.yaml```

# Create namespace
```kubectl create namespace qa```

# Delete a resource
```kubectl -n argocd  delete configmap/github-cm```

# Get a resource
kubectl -n argocd get ApplicationSet

# Describe a resource
kubectl -n argocd describe applicationset/calculator-services-dev

# Port forward
kubectl -n dev port-forward deploy/multiple-operation-svc-go-kit-deploy 8080:8080