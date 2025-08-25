- sec2
  - [gihyo-ms-dev-book/handson/sec2/2.2.5-messagebroker/order | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.5-messagebroker/order)


# RabbitMQの起動

```bash
docker run -it --rm --name rabbitmq --network $DOCKER_NETWORK rabbitmq:3.11-management
```

# プロジェクト作成

```bash
cd $PROJECT_DIR/application/order

go mod init gihyo/order

go get github.com/rabbitmq/amqp091-go
```

# 実行

```bash
(
  cd $PROJECT_DIR/application/order &&
  go run main.go
)
# 2025/08/25 17:39:11  [x] Sent {"ID":"test_id","CustomerId":"test","CustomerName":"customer name","OrderItem":[]}
```