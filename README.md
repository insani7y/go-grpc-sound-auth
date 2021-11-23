```migrate -path migrations -database "postgres://127.0.0.1/postgres?sslmode=disable&user=postgres&password=pass" up```

```protoc --go_out=plugins=grpc:internal/app internal/app/schema.proto```
