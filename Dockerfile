FROM golang:1.16
RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgres-client

# make wait-for-postgres.sh executable

RUN chmod +x wait-for-postgres.sh

# build go app

RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]