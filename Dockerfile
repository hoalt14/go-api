FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o myapp .

FROM scratch
COPY --from=builder /app/myapp /myapp
EXPOSE 8080
ENTRYPOINT ["/myapp"]