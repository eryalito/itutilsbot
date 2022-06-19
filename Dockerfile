FROM docker.io/library/golang:1.18 as builder

WORKDIR /go/src/github.com/eryalus/itutilsbot

COPY go.mod go.sum main.go ./
COPY commands/ commands/
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o itutilsbot main.go

FROM scratch 

WORKDIR /root

COPY --from=builder /go/src/github.com/eryalus/itutilsbot/itutilsbot ./

ENTRYPOINT [ "/root/itutilsbot" ]