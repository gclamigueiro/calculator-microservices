
# Testing Services

## Sum services   

```kubectl -n dev port-forward service/sum-svc-gokit-svc 8080:8082```

```kubectl -n dev port-forward sum-svc-gokit-deploy-77ff458858-z6nr8 8082:8080```

```curl -XPOST -d'{"a":5, "b": 3}' localhost:8082/```