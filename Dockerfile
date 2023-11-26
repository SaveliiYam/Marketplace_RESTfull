FROM golang:1.21.4-bullseye

RUN go version
ENV GOPATH=/

COPY ./ ./

#install postgres
RUN apt-get update
RUN apt-get -y install postgresql-client

#make wait-for-postgres.sh
RUN chmod +x wait_for_postgres.sh

#build go app
RUN go mod download
RUN go build -o market-app ./cmd/main.go