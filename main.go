package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./stations"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//defina a default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	var m stations.Stations
	err := json.Unmarshal(msg.Payload(), &m)
	if err != nil {
		fmt.Print(err)
	}

	for i := 0; i < len(m.Stations); i++ {
		fmt.Printf("SSID: %s\n", m.Stations[i].Ssid)
		fmt.Printf("BSSID: %s\n", m.Stations[i].Bssid)
		fmt.Printf("Signal: %d db\n\n", m.Stations[i].Signal)
	}

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
