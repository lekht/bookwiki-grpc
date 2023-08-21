FROM golang:latest as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o app ./cmd/app/main.go

FROM ubuntu AS production

WORKDIR /app

RUN apt-get update
RUN apt-get -y install mysql-client

COPY --from=builder /app/app ./
COPY --from=builder /app/wait-for-mysql.sh ./
COPY --from=builder /app/.env ./

RUN chmod +x wait-for-mysql.sh