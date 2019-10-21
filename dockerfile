# build stage
FROM golang:alpine as builder
RUN apk add --no-cache git
WORKDIR /github.com/IamStubborN/test
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go mod download
RUN go build -o test ./cmd

# final stage
FROM alpine:latest
WORKDIR /root/test/
COPY --from=builder /github.com/IamStubborN/test .

ENTRYPOINT ["./test"]