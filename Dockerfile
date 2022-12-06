FROM golang:1.19 AS builder
WORKDIR /go/src/github.com/fdonzello/go-with-nats/
COPY . .
ARG app=none
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o appx cmd/${app}/main.go

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/fdonzello/go-with-nats/appx .

CMD ["./appx"]