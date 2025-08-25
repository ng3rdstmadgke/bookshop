package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// エラーをチェックするヘルパー関数
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// RabbitMQサーバーに接続
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// チャンネルを開く
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// キューを宣言
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

	// キューからメッセージを読み取る
	// https://pkg.go.dev/github.com/rabbitmq/amqp091-go@v1.10.0#Channel.Consume
	msgs, err := ch.Consume(
		q.Name, // queueu string キューの名前
		"",     // consumer string コンシューマを識別するためのID(空文字列で自動生成)
		true,   // autoAck bool Consumerがメッセージを受け取ったら自動的にACKを送信 (受信した瞬間にメッセージがキューから削除される)
		false,  // exclusive bool 複数のコンシューマが同じキューを利用できる
		false,  // noLocal bool 自分が送信したメッセージも受け取れる
		false,  // noWait bool サーバーからの応答を待つ
		nil,    // args Table 引数(オプション)
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{} // struct{} 型のチャンネルを宣言 (値はnil、truct{}は空の構造体でメモリを使わない)

	// ゴルーチンでメッセージを処理
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever // nil チャンネルから読み取りを試行 (永遠にブロックされるのでプログラムが終了しない)
}
