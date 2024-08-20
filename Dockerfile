# Stage 1: Build
FROM golang:1.23.0-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o myapp .

# Stage 2: Deploy
FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/myapp /myapp

EXPOSE 8080
ENTRYPOINT ["/myapp"]