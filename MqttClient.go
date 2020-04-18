package main

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func createMQTTClient(clientID string, mqttinfo string, handler MQTT.MessageHandler, subScribeList []string) {
	opts := MQTT.NewClientOptions().AddBroker(mqttinfo)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(handler)
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, v := range subScribeList {
		if token := c.Subscribe(v, 0, nil); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
}
