package services

import (
	"context"
	"encoding/json"
	"fmt"

	"estudo.com/dtos"
	"estudo.com/helpers"
	"github.com/segmentio/kafka-go"
)

var kafkaReader *kafka.Reader

func GetKafkaContent() string {
	var conn = helpers.GetKafkaConnection()
	getKafkaReader(conn)
	var message = readMessage()
	return message
}

func getKafkaReader(conn helpers.KafkaConnection) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{conn.Broker},
		Topic:   conn.Topic,
		GroupID: conn.Group,
	})
}

func readMessage() string {
	ctx := context.Background()
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		var dto dtos.TransactionDTO
		json.Unmarshal(msg.Value, &dto)
		fmt.Println("name: ", dto.Name)
		fmt.Println("value: ", dto.Value)
	}
}
