# 起動方法

gRPCサーバーの起動

```bash
(
  cd $PROJECT_DIR/application/catalogue &&
  go run main.go
)
```

GraphQLサーバーの起動

```bash
(
  cd $PROJECT_DIR/application/bff &&
  npm start
)
```

Reactサーバーの起動

```bash
(
  cd $PROJECT_DIR/application/frontend &&
  npm start
)
```

- gRPCサーバー: http://localhost:50051
- GraphQLサーバー: http://localhost:4000
- Reactサーバー: http://localhost:3000