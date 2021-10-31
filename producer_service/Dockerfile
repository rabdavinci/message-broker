FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/github.com/rabdavinci/message-broker/producer_service
WORKDIR $GOPATH/src/github.com/rabdavinci/message-broker/producer_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/producer_service /

FROM alpine
COPY --from=builder ur_go_user_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent
ENTRYPOINT ["/producer_service"]
