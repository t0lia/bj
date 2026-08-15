FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bj ./cmd


FROM scratch

COPY --from=builder /app/bj /bj

EXPOSE 8080

ENTRYPOINT ["/bj"]
