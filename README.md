# Go API

## Local

### Update Module

```shell
go mod tidy
```

### Build Execute File

```shell
go build -o test
```

### Run test

```shell
./test
```

## Docker

### Build Image

```shell
docker build -t go-api .
```

### Run Container

```shell
docker run -d --name test -p 8080:8080 go-api
```

## Access

[http://localhost:8080/](http://localhost:8080/)
