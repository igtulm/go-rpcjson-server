FROM golang:1.11.5

RUN apt-get update && apt-get install -y postgresql-client && apt-get clean

RUN go get github.com/kelseyhightower/envconfig && \
    go get github.com/lib/pq && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/rpc && \
    go get github.com/gorilla/rpc/json && \
    go get github.com/stretchr/testify/require
