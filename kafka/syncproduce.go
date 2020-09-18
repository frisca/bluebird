package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type KafkaProducer struct {
	Connection sarama.SyncProducer
	ErrRes     error
	KeyEncoder sarama.Encoder
	Value      sarama.Encoder
	Topic      string
	Done       bool
}

func (kafkareq *KafkaProducer) Process() {
	kafkareq.Done = kafkareq.Send()
}

func (kafka *KafkaProducer) Send() bool {
	producer := kafka.Connection

	producerMsg := &sarama.ProducerMessage{
		Topic: kafka.Topic,
		Key:   kafka.KeyEncoder,
		Value: kafka.Value,
	}

	err := sendProducer(producer, producerMsg)
	if err != nil {
		fmt.Println("Error :", err)
		return false
	}

	return true
}

func sendProducer(producer sarama.SyncProducer, produceMessage *sarama.ProducerMessage) error {
	defer func() {
		if err := producer.Close(); err != nil {
			logs.Error("Producer Can't Close : ", err)
		}
	}()

	_, _, err := producer.SendMessage(produceMessage)
	if err != nil {
		logs.Error("Producer Send Message : ", err)
		return err
	}
	return nil

}
