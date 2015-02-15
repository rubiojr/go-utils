package main

import (
	"errors"
	"fmt"
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"net/url"
	"os"
)

// tcp://user:password@host:port
func pushMsg(clientId, brokerUrl, topic, msg string) (bool, error) {

	uri, _ := url.Parse(brokerUrl)
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	user := uri.User.Username()
	if user != "" {
		opts.SetUsername(uri.User.Username())
		password, _ := uri.User.Password()
		if password != "" {
			opts.SetPassword(password)
		}
	}

	opts.SetClientId(clientId)

	client := mqtt.NewClient(opts)
	_, err := client.Start()
	if err != nil {
		return false, errors.New("Error starting the MQTT Client: " + err.Error())
	}

	<-client.Publish(0, topic, msg)

	return true, nil
}

func main() {
	pushMsg("pushmtr", os.Getenv("MQTT_URL"), "topic", "message")
}
