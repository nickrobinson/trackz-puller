package main

import (
	"encoding/json"
	"log"
	"os"
	"flag"

	"github.com/nickrobinson/trackz-puller/stations"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strconv"
)

//define a default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	var m stations.Stations
	err := json.Unmarshal(msg.Payload(), &m)
	if err != nil {
		log.Print(err)
	}

	for i := 0; i < len(m.Stations); i++ {
		log.Printf("SSID: %s\n", m.Stations[i].Ssid)
		log.Printf("BSSID: %s\n", m.Stations[i].Bssid)
		log.Printf("Signal: %d db\n\n", m.Stations[i].Signal)
	}

}

func main() {
	var mqttServer = flag.String("server", "test.mosquitto.org", "MQTT Server to connect to")
	var mqttPort = flag.Int("port", 1883, "MQTT Server Port")
	var mqttClientId = flag.String("client", "testgoid", "MQTT Client Identifier")
	flag.Parse()

	log.SetOutput(os.Stderr)
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + *mqttServer + ":" + strconv.Itoa(*mqttPort))
	opts.SetClientID(*mqttClientId)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("Client connected")

	if token := client.Subscribe("/nickrobi/0001/aps", byte(0), f); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

	for {
	}

}
