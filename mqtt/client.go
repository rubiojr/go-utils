package main

import (
	"errors"
	"fmt"
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"net/url"
)

// tcp://user:password@host:port
func PushMsg(clientId, brokerUrl, topic, msg string) error {

	if brokerUrl == "" {
		panic("Invalid broker URL")
	}

	uri, _ := url.Parse(brokerUrl)
	if uri.Scheme != "tcp" {
		panic("Invalid broker URL scheme")
	}

	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	if uri.User != nil {
		user := uri.User.Username()
		opts.SetUsername(user)
		password, _ := uri.User.Password()
		if password != "" {
			opts.SetPassword(password)
		}
	}

	opts.SetClientId(clientId)

	client := mqtt.NewClient(opts)
	_, err := client.Start()
	if err != nil {
		return errors.New("Error starting the MQTT Client: " + err.Error())
	}

	<-client.Publish(0, topic, msg)

	return nil
}
