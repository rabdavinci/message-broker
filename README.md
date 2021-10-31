# Message broker GRPC

This repo contains solutions of "Message broker" on Golang

## DESCRIPTION

Write a simple message broker, SDK for it in the form of a library and examples of client applications (producer, consumer).
Operation protocol - gRPC

## How 2RUN

1. Clone project `git clone git@github.com:rabdavinci/message-broker . `
2. Run postgres and create tables from each service/migrations
3. Run broker service `go run broker_service/main.go`
4. Run producer service `go run producer_service/main.go`
5. Run user service `go run user_service/main.go`

## TODO

1. Finish Dockerize
2. Finish CI/CD
3. Add Scalability
4. Add documentation
