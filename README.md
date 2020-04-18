# mqtt2kafka
binding mqtt topic and redirect to kafka server.

## Usage

1. Open `main.go`, modify `MQTTConfig`, `KafkaConfig`.
2. Open `main.go`, in `createMQTTClient`, you have to overwrite the topic want to be listen in `[]string{"#"}`. (it can be many).
3. run `go run main.go MqttClient.go KafkaClient.go`
