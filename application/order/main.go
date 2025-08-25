package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderItem struct {
	ItemId    int `bson:"itemId"`
	Quantity  int `bson:"quantity"`
	UnitPrice int `bson:"unitPrice"`
}

type OrderEvent struct {
	ID           string
	CustomerId   string
	CustomerName string
	OrderItem    []OrderItem
}

// エラーをチェックするヘルパー関数
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// RabbitMQサーバーへの接続
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// チャンネルの作成
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// キューの宣言
	// https://pkg.go.dev/github.com/rabbitmq/amqp091-go@v1.10.0#Channel.QueueDeclare
	q, err := ch.QueueDeclare(
		"order", // name string キューの名前
		false,   // durable サーバーが落ちるとキューも消える
		false,   // autoDelete Consumerがいなくてもキューは残り続ける
		false,   // exclusive 複数接続から利用可能
		false,   // noWait bool サーバーからの応答を待つ
		nil,     // args Table 引数(オプション)
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 送信するメッセージの作成
	orderEvent := OrderEvent{
		ID:           "test_id",
		CustomerId:   "test",
		CustomerName: "customer name",
		OrderItem:    []OrderItem{},
	}

	body, err2 := json.Marshal(orderEvent)
	failOnError(err2, "Failed to marshal")

	// メッセージをキューにパブリッシュ
	// https://pkg.go.dev/github.com/rabbitmq/amqp091-go@v1.10.0#Channel.PublishWithContext
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key (キュー名)
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
