FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Adiciona a flag CGO_ENABLED=0 e define a arquitetura correta
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

#RUN go build -o main .

FROM alpine:latest  

WORKDIR /app

COPY --from=builder /app/main .

# Permissão de execução (caso necessário)
RUN chmod +x /app/main

CMD ["/app/main"]
