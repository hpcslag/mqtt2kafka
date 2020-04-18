package main

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

type Server struct {
	AccessLogProducer sarama.AsyncProducer
}

// 初始化Kafka Producer的相關設定
func newAccessLogProducer(brokerList []string) sarama.AsyncProducer {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write access log entry:", err)
		}
	}()

	return producer
}

func createKafkaClient(brokers *string) *Server {
	brokerList := strings.Split(*brokers, ",")
	log.Printf("Kafka brokers: %s", strings.Join(brokerList, ", "))

	kafka = &Server{
		AccessLogProducer: newAccessLogProducer(brokerList),
	}

	return kafka
}

func (kafka *Server) sendToKafka(kafkaTargetTopic string, raw string) {
	kafka.AccessLogProducer.Input() <- &sarama.ProducerMessage{
		Topic: kafkaTargetTopic,
		Value: sarama.StringEncoder(raw),
	}
}
