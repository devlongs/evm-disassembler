FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o evm-disassembler ./cmd/server

FROM alpine:3.17
WORKDIR /app


COPY --from=builder /app/evm-disassembler /usr/local/bin/evm-disassembler


COPY configs/ /app/configs/

EXPOSE 8080

ENTRYPOINT ["evm-disassembler"]
