# Message broker GRPC

This repo contains solutions of "Message broker" on Golang

## DESCRIPTION

Write a simple message broker, SDK for it in the form of a library and examples of client applications (producer, consumer).
Operation protocol - gRPC

## How 2RUN

1. Clone project `git clone git@github.com:rabdavinci/message-broker . `
2. Run postgres and create tables from each service/migrations <br />
`docker run --name urecruitdb -e POSTGRES_PASSWORD='qwerty' -p 5430:5432 postgres`
4. Run broker service `go run broker_service/main.go`
5. Run producer service `go run producer_service/main.go`
6. Run user service `go run user_service/main.go`

## TODO

1. Finish Dockerize
2. Finish CI/CD
3. Add Scalability
4. Add documentation

## SCREENSHOTS
![image](https://user-images.githubusercontent.com/30826165/139599282-c436c3f6-6c3b-4328-a923-9b00b1d68f0e.png)
