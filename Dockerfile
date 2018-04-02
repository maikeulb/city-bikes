FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/maikeulb/city-bikes

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o city-bikes -a -installsuffix cgo main.go app.go model.go


FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/maikeulb/city-bikes .

CMD ["./city-bikes"]
