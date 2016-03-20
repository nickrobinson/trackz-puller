package main

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//defina a default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://mqtt.isengard.io:1883")
	opts.SetClientID("testgoid")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/nickrobi/0001/aps", byte(0), f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
	}

}
