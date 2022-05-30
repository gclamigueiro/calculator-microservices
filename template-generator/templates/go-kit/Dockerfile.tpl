FROM golang:alpine AS build
LABEL maintainer=gclamigueiro

WORKDIR /app
COPY ./services/{{.APIName}} .
RUN go build ./cmd/main.go

ENTRYPOINT [ "/app/main" ][]

# to create image
# docker build -t {{.APINamespace}}{{.APIName}} .
# to run image
# docker run -d -p 8080:8080 {{.APINamespace}}{{.APIName}}