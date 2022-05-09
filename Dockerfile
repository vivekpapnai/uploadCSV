FROM golang:1.16.6-alpine3.14 AS builder
# Install Certificate
RUN apk add -U --no-cache ca-certificates

FROM scratch AS app

# Copy Certificate
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

WORKDIR /server
ENV GO111MODULE=on
COPY go.mod /server/
COPY go.sum /server/

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o /go/bin/uploadCSV

FROM scratch as app
COPY --from=builder /go/bin/uploadCSV /go/bin/uploadCSV
EXPOSE 8080
ENTRYPOINT ["/go/bin/uploadCSV"]