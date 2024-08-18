FROM golang:1.23.0-alpine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/go-app .

EXPOSE 8080

CMD ["./out/go-app"]