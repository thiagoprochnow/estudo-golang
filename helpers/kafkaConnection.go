package helpers

type KafkaConnection struct {
	Broker string
	Topic  string
	Group  string
}

func GetKafkaConnection() KafkaConnection {
	return KafkaConnection{
		Broker: "localhost:29092",
		Topic:  "topico_teste",
		Group:  "grupo_teste",
	}
}
