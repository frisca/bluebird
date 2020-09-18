package service

import (
	"errors"
	"bluebird/kafka"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)


type ServicePush struct {
	Topic       string
	Brokerlist  string
	ConfigKafka *sarama.Config
}

func (s *ServicePush) PushToKafka(bdata []byte) error {
	logs.Info("Produce to Kafka")
	kafkaProducer := kafka.GetConectKafka(s.Brokerlist, s.ConfigKafka)
	if kafkaProducer.ErrRes != nil {
		logs.Error("Err Client :", kafkaProducer.ErrRes)
		return kafkaProducer.ErrRes
	}
	
	kafkaProducer.Topic = s.Topic
	kafkaProducer.Value = sarama.ByteEncoder(bdata)
	kafkaProducer.Process()
	if kafkaProducer.Done == false {
		return errors.New("Sending to Kafka Failed")
	}
	return nil
}

func (s *ServicePush) CheckConnectKafka() (bool, error) {
	kafkaProducer := kafka.GetConectKafka(s.Brokerlist, s.ConfigKafka)
	if kafkaProducer.ErrRes != nil {
		return false, kafkaProducer.ErrRes
	}
	return true, kafkaProducer.ErrRes
}
