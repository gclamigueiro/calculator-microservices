# multiple-operation-svc-go-kit Microservice

the service recibe an array of operations and call the correspondent sevice

## Dependencies
- sum-svc-gokit-svc
- subtract-svc-gokit-svc

## Test de Service

### Port Forward to dependant services
kubectl -n dev port-forward service/sum-svc-gokit-svc 8081:8080
kubectl -n dev port-forward service/subtract-svc-gokit-svc 8082:8080

### Run the service 
```
go run  .\cmd\main.go
```

### example request 

curl -XPOST -d'{"operations":[{"param1":5, "param2": 3, "operation":"+"},{"param1":5, "param2": 3, "operation":"+"} ]}' localhost:8080/v1/calculator/service

