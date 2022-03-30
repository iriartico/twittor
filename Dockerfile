# STATICS
# FROM golang:1.18 AS builder
# RUN apt-get update
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64
# WORKDIR /go/src/app
# COPY go.mod .
# RUN go mod download
# COPY . .
# RUN go install

# FROM scratch
# COPY --from=builder /go/bin/app .
# ENTRYPOINT [ "./main" ]


# FROM WEB
FROM golang:1.18 AS builder
RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/github.com/iriartico/twittor
COPY go.mod .
RUN go mod download
COPY . .
RUN go build main.go

# FROM scratch
# COPY --from=builder /go/src/github.com/iriartico/twittor/.env .
# COPY --from=builder /go/src/github.com/iriartico/twittor/main .
# ENTRYPOINT [ "./main" ]


# FROM golang:1.18 AS builder
# WORKDIR /go/src/github.com/iriartico/twittor
# # RUN go get -d -v -a
# COPY main.go .
# RUN  CGO_ENABLED=0 GOOS=linux go build -a \
#      install suffix cgo -o main .

FROM alpine:3.15
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY .env .
COPY --from=builder /go/src/github.com/iriartico/twittor/main .
ENTRYPOINT [ "./main" ]