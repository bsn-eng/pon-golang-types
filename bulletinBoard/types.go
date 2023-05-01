package bulletinboard

import (
	"time"

	pahoMQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

var (
	TCP = "tcp"
	WSS = "wss"
)

var (
	RelayPublishLimit = time.Duration(100)
)

type MQTTTopic string

type RelayMQTTOpts struct {
	Broker   string
	Port     uint64
	ClientID string
	UserName string
	Password string
}

type relayMQTT struct {
	Broker string

	ClientOptions *pahoMQTT.ClientOptions
	Client        pahoMQTT.Client

	log *logrus.Entry
}
