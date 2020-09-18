package services


import (
	"encoding/json"
	"bluebird/kafka"
	"bluebird/kafka/service"

	"github.com/astaxie/beego/logs"
	"fmt"
)

type KafkaService struct {
	success bool
}

func (addLog *KafkaService) AddToKafka(topic string, req interface{}) (err error) {
	procedure := kafka.GetConnKakfaProcedure()
	var kafkaTopic string
	if procedure.ErrRes != nil {
		logs.Error("Kafka Connect Error ", procedure.ErrRes.Error())
	}
	logs.Info("Kafka connected")

	kafkaTopic = topic
	kafka := service.ServicePush{
		Topic:       kafkaTopic,
		Brokerlist:  kafka.GetKafkaBroker(),
		ConfigKafka: kafka.GetConfigKafka(),
	}
	fmt.Println("kafka: ", kafka)
	formInBytes, _ := json.Marshal(req)
	errKafka := kafka.PushToKafka(formInBytes)

	addLog.success = true
	return errKafka
}

func (addLog *KafkaService) GetStatus() bool {
	return addLog.success
}