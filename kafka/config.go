package kafka

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
)

var (
	kafkaBrokerUrl     string
	kafkaTopicBlueBird    string
	kafkaClient        string
	kafkaTimeout       int
	maxmsgbyte         int
)

func init() {
	kafkaBrokerUrl = beego.AppConfig.DefaultString("kafka.brokers", "localhost:9092")
	kafkaTopicBlueBird = beego.AppConfig.DefaultString("kafka.topics.bluebird", "order-topic")
	kafkaTimeout = beego.AppConfig.DefaultInt("kafka.produce.timeout", 10)
	maxmsgbyte = beego.AppConfig.DefaultInt("kafka.produce.maxmsgbyte", 50000000)
}

func GetConectKafka(brokerList string, config *sarama.Config) KafkaProducer {
	producer, err := sarama.NewSyncProducer(strings.Split(brokerList, ","), config)
	return KafkaProducer{
		Connection: producer,
		ErrRes:     err,
	}
}

func GetKafkaBroker() string {
	return kafkaBrokerUrl
}

func GetTopicBlueBird() string {
	return kafkaTopicBlueBird
}

func GetConnKakfaProcedure() KafkaProducer {
	producer, err := sarama.NewSyncProducer(strings.Split(GetKafkaBroker(), ","), GetConfigKafka())
	return KafkaProducer{
		Connection: producer,
		ErrRes:     err,
	}
}

func GetConfigKafka() *sarama.Config {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	configKafka := sarama.NewConfig()
	configKafka.Producer.Return.Successes = true
	configKafka.Producer.Partitioner = sarama.NewRandomPartitioner
	configKafka.Producer.MaxMessageBytes = maxmsgbyte
	configKafka.Producer.Timeout = 5 * time.Second
	configKafka.Net.DialTimeout = 2 * time.Second
	configKafka.Net.ReadTimeout = 5 * time.Second
	configKafka.Net.WriteTimeout = 5 * time.Second

	return configKafka
}
