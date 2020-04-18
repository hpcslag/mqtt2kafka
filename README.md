# mqtt2kafka
binding mqtt topic and redirect to kafka server.

## Usage

1. Open `main.go`, modify `MQTTConfig`, `KafkaConfig`.
2. Open `main.go`, in `createMQTTClient`, you have to overwrite the topic want to be listen in `[]string{"#"}`. (it can be many).
3. run `go run main.go MqttClient.go KafkaClient.go`

## With Password

Modify `MqttClient`:
```
opts.SetUsername(*username)
opts.SetPassword(*password)
```

## Handler

`main.go: mqttHandler`:
```go
func mqttHandler(client MQTT.Client, msg MQTT.Message) {
  topic := msg.Topic()
  payload := msg.Payload()

  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
  
  //custom...
  mqttData := &mqttModel{}
  json.Unmarshal([]byte(payload), mqttData)
}
```
