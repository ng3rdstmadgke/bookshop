- [gihyo-ms-dev-book/handson/sec2/2.2.1-backend/catalogue | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.1-backend/catalogue)

# プロジェクト作成

```bash
cd $PROJECT_DIR/application/catalogue

go mod init gihyo/catalogue

go get google.golang.org/grpc
```

# gRPCサービスのコンパイル


```bash
(
  cd $PROJECT_DIR/application/catalogue/proto/book && \
  protoc \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    catalogue.proto
)
```

# 起動

```bash
(
  cd $PROJECT_DIR/application/catalogue &&
  go run main.go
)
```


```bash
# 公開されているサービスの確認
grpcurl -plaintext localhost:50051 list
# book.Catalogue
# grpc.reflection.v1.ServerReflection
# grpc.reflection.v1alpha.ServerReflection

# book.Catalogueサービスのメソッドを確認
grpcurl -plaintext localhost:50051 list book.Catalogue
# book.Catalogue.GetBook

# book.Catalogue.GetBookメソッドにアクセス
grpcurl -plaintext -d '{"id": 1}' localhost:50051 book.Catalogue.GetBook
# {
#   "book": {
#     "id": 1,
#     "title": "The Awakening",
#     "author": "Kate Chopin",
#     "price": 1000
#   }
# }
```