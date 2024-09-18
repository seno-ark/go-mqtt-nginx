FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-mqtt

FROM alpine:latest
COPY --from=builder /app/go-mqtt .

EXPOSE 1883
EXPOSE 1884

ENTRYPOINT ["./go-mqtt"]