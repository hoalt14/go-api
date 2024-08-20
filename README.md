# Go API

## Quickstart

```shell
docker pull hoalt14/go-api:v1
```

```shell
docker run -d --name test -p 8080:8080 hoalt14/go-api:v1
```

## Testing

[http://localhost:8080/](http://localhost:8080/)

```shell
curl localhost:8080
curl localhost:8080/hello
curl localhost:8080/health

# High Load CPU
curl 'localhost:8080/fib?n=40'
```
