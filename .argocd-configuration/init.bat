kubectl apply -f .\namespaces.yaml
kubectl apply -f .\github-secret.yaml
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f .\appprojects\app-project-calculator-dev.yaml
:: kubectl apply -f .\appprojects\app-project-calculator-qa
kubectl apply -f .\applicationset-develop.yaml

