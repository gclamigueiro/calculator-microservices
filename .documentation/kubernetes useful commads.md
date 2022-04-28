# Deploy a resource 
```kubectl apply -f .\app-project-calculator-dev.yaml```

# Create namespace
```kubectl create namespace qa```

# Delete a resource
```kubectl -n argocd  delete configmap/github-cm```

# Get a resource
kubectl -n argocd get ApplicationSet