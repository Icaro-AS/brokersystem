package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka" 

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topic     []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topic:     topics,
	}
}

func(c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(c.Topic, nil)

	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil{
			msgChan <- msg
		}
	}
	
}