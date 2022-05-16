FROM golang:alpine AS build
LABEL maintainer=gclamigueiro

WORKDIR /app
COPY ./services/multiple-operation-svc-go-kit .
RUN go build ./cmd/main.go

ENTRYPOINT [ "/app/main" ][]

# to create image
# docker build -t multiple-operation-svc-go-kit .
# to run image
# docker run -d -p 8080:8080 multiple-operation-svc-go-kit