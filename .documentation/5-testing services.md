
# Testing Services

## Sum services   


```kubectl -n dev port-forward  deployment/sum-svc-gokit-deploy 8080:8082```

```kubectl -n dev port-forward service/sum-svc-gokit-svc 8080:8082```

```kubectl -n dev port-forward sum-svc-gokit-deploy-5dcfc96f98-hgw68 8082:8080```

kubectl port-forward deployment/sum-svc-gokit-svc-deploy 8080:8082

```curl -XPOST -d'{"a":5, "b": 3}' localhost:8082/```