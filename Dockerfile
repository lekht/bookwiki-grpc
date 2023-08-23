FROM golang:latest as builder
WORKDIR /service

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go test ./...

RUN go build -o service ./cmd/service/service.go

FROM ubuntu AS production

WORKDIR /service

RUN apt-get update
RUN apt-get -y install mysql-client

COPY --from=builder /service/service ./
COPY --from=builder /service/wait-for-mysql.sh ./
COPY --from=builder /service/.env ./

RUN chmod +x wait-for-mysql.sh