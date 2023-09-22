package main

import (
	"fmt"
	"mqtt_stuff/pkg/mqtt"
	"time"

	paho_mqtt "github.com/eclipse/paho.mqtt.golang"
)

const resourceName = "myResource"

const brokerAddress = ":1883"

func main() {
	client := mqtt.New()

	subscribeToTopics(client, brokerAddress)

	for {
		publishHeartbeat(client)
		time.Sleep(5 * time.Second)
	}
}

func publishAddUser(client mqtt.MQTTServer, rfidTag string) {
	client.Publish(brokerAddress, "frontdoor/adduser", rfidTag)
}

func publishRemoveUser(client mqtt.MQTTServer, rfidTag string) {
	client.Publish(brokerAddress, resourceName+"/removeuser", rfidTag)
}

func publishHeartbeat(client mqtt.MQTTServer) {
	client.Publish(brokerAddress, resourceName+"/heartbeat", time.Now().Unix())
}

func subscribeToTopics(client mqtt.MQTTServer, address string) {
	client.Subscribe(address, resourceName+"/adduser", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
		rfidTag := string(msg.Payload())
		fmt.Printf("Add user with RFID tag: %s\n", rfidTag)
		// store user to memory
		// could respond with something -- probably not necessary though
	})

	client.Subscribe(address, resourceName+"/removeuser", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
		rfidTag := string(msg.Payload())
		fmt.Printf("Delete user with RFID tag: %s\n", rfidTag)
		// delete user from memory
		// could respond with something -- probably not necessary though
	})

	client.Subscribe(address, resourceName+"/request_verify", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
		rfidTag := string(msg.Payload())
		fmt.Printf("Verification request for RFID tag: %s\n", rfidTag)
		client.Publish(brokerAddress, resourceName+"/respond_verify", "1234567890abcdef")
	})

	client.Subscribe(address, resourceName+"/request_hash", func(c paho_mqtt.Client, msg paho_mqtt.Message) {
		fmt.Printf("hash of entire rfid list: %s\n", "")
		client.Publish(brokerAddress, resourceName+"/respond_hash", "1234567890abcdef")
	})
}
