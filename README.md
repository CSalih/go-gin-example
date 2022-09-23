# go-gin-example

Example RESTful API written in Go using Gin framework.


## Getting Started

```bash
go run cmd/example-gin/main.go
```

### Get user
```bash
curl --location --request GET 'http://localhost:8080/api/v1/users/foo'
```

### Update user
```bash
curl --location --request PUT 'http://localhost:8080/api/v1/users/foo' \
--header 'Authorization: Basic bWFudToxMjM=' \
--header 'Content-Type: application/json' \
--data-raw '{
  "status": "Hello World"
}'
```