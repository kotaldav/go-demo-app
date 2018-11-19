FROM golang:onbuild as builder

WORKDIR /go/src/app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#---

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/src/app/main .

CMD ["./main"]
