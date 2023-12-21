FROM golang:1.21.5

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o jwt

EXPOSE 9876

CMD ["./jwt"]
