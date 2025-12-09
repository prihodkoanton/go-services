FROM golang:1.25 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./services/user-service/cmd

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 9998

CMD ["./user-service"]
