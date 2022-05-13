FROM golang:1.16.6-alpine3.14 AS builder
#RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
#
#RUN apk add -U --no-cache ca-certificates

#COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /go/bin/ssl/certs/

WORKDIR /server
ENV GO111MODULE=on
COPY go.mod /server/
COPY go.sum /server/

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o /go/bin/uploadCSV

FROM scratch
COPY --from=builder /go/bin/uploadCSV /go/bin/uploadCSV
EXPOSE 8080
ENTRYPOINT ["/go/bin/uploadCSV"]
