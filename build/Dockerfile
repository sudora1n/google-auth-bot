FROM golang:1.20 as builder

WORKDIR /app

COPY . ./ 
RUN go mod download
RUN go build -v -o server ./cmd/microservice-api/main.go


FROM debian:trixie-slim

WORKDIR /app
COPY --from=builder /app/server /app/server

EXPOSE 8080
ENTRYPOINT [ "/app/server" ]