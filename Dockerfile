FROM golang:1.15.6-alpine

WORKDIR /app/go-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/go-app .

EXPOSE 8080

CMD ["./out/go-app"]