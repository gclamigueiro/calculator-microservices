FROM golang:alpine AS build
LABEL maintainer=gclamigueiro

WORKDIR /app
COPY . .
RUN go build ./cmd/main.go

ENTRYPOINT [ "/app/main" ][]

# to create image
# docker build -t sum-svc-gokit .
# to run image
# docker run -d -p 8080:8080 sum-svc-gokit