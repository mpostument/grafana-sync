# Build App
FROM golang:1.17.11-alpine AS builder

WORKDIR ${GOPATH}/src/github.com/mpostument/grafana-sync
COPY . ${GOPATH}/src/github.com/mpostument/grafana-sync

RUN go build -o /go/bin/grafana-sync .


# Create small image with binary
FROM alpine:3.15

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/grafana-sync /usr/bin/grafana-sync

ENTRYPOINT ["/usr/bin/grafana-sync"]
