FROM golang:alpine AS build
LABEL maintainer=gclamigueiro

WORKDIR /app
COPY ./services/github.com/gclamigueiro/sum-svc-gokit .
RUN go build ./cmd/main.go

ENTRYPOINT [ "/app/main" ][]

# to create image
# docker build -t github.com/gclamigueiro/sum-svc-gokit .
# to run image
# docker run -d -p 8080:8080 github.com/gclamigueiro/sum-svc-gokit