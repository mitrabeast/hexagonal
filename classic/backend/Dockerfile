FROM golang:1.19

WORKDIR /go/src/backend
COPY . .

RUN go mod download -x
RUN go build -o backend

CMD ["./backend", "/etc/config.json"]