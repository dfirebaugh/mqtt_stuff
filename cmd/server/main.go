package main

import (
	"mqtt_stuff/pkg/mqtt"
	"time"

	paho_mqtt "github.com/eclipse/paho.mqtt.golang"
)

const resourceName = "myResource"

const brokerAddress = "tcp://localhost:1883"

func main() {
	client := mqtt.New()

	// Subscribing to topics
	subscribeToTopics(client, brokerAddress)

	for {
		time.Sleep(5 * time.Second)
	}
}

func publishAddUser(client mqtt.MQTTServer, rfidTag string) {
	client.Publish(brokerAddress, "frontdoor/adduser", rfidTag)
}

func publishRemoveUser(client mqtt.MQTTServer, rfidTag string) {
	client.Publish(brokerAddress, resourceName+"/removeuser", rfidTag)
}

func publishRequestHash(client mqtt.MQTTServer, rfidTag string) {
	client.Publish(brokerAddress, resourceName+"/request_hash", rfidTag)
}

func subscribeToTopics(client mqtt.MQTTServer, address string) {
	client.Subscribe(address, resourceName+"/adduser", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
	})

	client.Subscribe(address, resourceName+"/respond_verify", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
	})

	client.Subscribe(address, resourceName+"/respond_hash", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
	})
}
