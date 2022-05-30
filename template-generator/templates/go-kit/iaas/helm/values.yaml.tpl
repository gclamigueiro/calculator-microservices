name: {{.APIName}}

image:
  name: usernameDocker/{{.APIName}}
  tag:  latest

config:
  port: 8080

ingress:
  path: "/v1/calculator/service"

