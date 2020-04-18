package main

import (
	"flag"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var MQTTConfig = map[string]string{
	"MQTT_HOST": "tcp://xxx.xxx.xxx.xxx",
	"ClientID":  "Test-Plugin",
}

var KafkaConfig = map[string]*string{
	"Brokers": flag.String("brokers", "localhost:9092", "TEST_BROKER"),
}

var kafka = &Server{} //set as global

func main() {

	//init kafka (many...)
	kafka = createKafkaClient(KafkaConfig["Brokers"])

	//init mqtt (many...)
	createMQTTClient(MQTTConfig["ClientID"], MQTTConfig["MQTT_HOST"], mqttHandler, []string{"#"})

	ch := make(chan int, 1)
	<-ch
}

func mqttHandler(client MQTT.Client, msg MQTT.Message) {
	kafka.sendToKafka("test7", "OK")
}
