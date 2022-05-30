## The chart API version
apiVersion: v1
## The name of the chart. 
name: {{.APINamespace}}{{.APIName}}
## A single-sentence description of this project
description: Calculator Service 
## A list of keywords about this project to quickly identify the chart's capabilities.
keywords:
- go
- go-kit
## The chart version, here set to `3.7.0`
version: 3.7.0
## List of maintainers, their names, and method of contact
maintainers:
- name: maintainer 
  email: maintainer@xyz.com