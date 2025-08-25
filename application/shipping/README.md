- sec2
  - [gihyo-ms-dev-book/handson/sec2/2.2.5-messagebroker/shipping | GitLab](https://gitlab.com/gihyo-ms-dev-book/handson/sec2/2.2.5-messagebroker/shipping)


# RabbitMQの起動

```bash
docker run -it --rm --name rabbitmq --network $DOCKER_NETWORK rabbitmq:3.11-management
```

# プロジェクト作成

```bash
cd $PROJECT_DIR/application/shipping

go mod init gihyo/shipping

go get github.com/rabbitmq/amqp091-go
```

# 実行

```bash
(
  cd $PROJECT_DIR/application/shipping &&
  go run main.go
)
```
