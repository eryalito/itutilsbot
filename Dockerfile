FROM docker.io/library/golang:1.18 as builder

WORKDIR /go/src/github.com/eryalus/itutilsbot

COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o itutilsbot cmd/main.go

FROM alpine:latest as certs

RUN apk update

FROM scratch 

WORKDIR /root

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/eryalus/itutilsbot/itutilsbot ./

ENTRYPOINT [ "/root/itutilsbot" ]