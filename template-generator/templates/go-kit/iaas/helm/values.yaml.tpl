name: {{.APIName}}

## define the image to execute with the Deployment 
image:
  tag: latest

config:
  port: 8080

path: "/v1/calculator/service"