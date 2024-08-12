FROM golang:1.22-alpine

ENV PORT=8080

WORKDIR /root/test

COPY go.mod main.go Makefile .

RUN go mod tidy

RUN go build -v -o test .

ENTRYPOINT ["./test"]
