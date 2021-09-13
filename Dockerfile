FROM golang:alpine3.14 as builder
WORKDIR /application
COPY go.mod ./
COPY go.sum ./
COPY src/ src/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o metrics src/main.go

FROM alpine:3.14.2 as job
WORKDIR /job
COPY --from=builder /application/metrics ./
RUN chmod 0755 metrics

ENTRYPOINT ["./metrics"]