package bulletinboard

import (
	"time"
)

var (
	TCP = "tcp"
	WSS = "wss"
)

var (
	RelayPublishLimit, _ = time.ParseDuration("15s")
)

type MQTTTopic string

type RelayMQTTOpts struct {
	Broker   string
	Port     uint64
	ClientID string
	UserName string
	Password string
}

type RelayHighestBid struct {
	Slot             uint64
	BuilderPublicKey string
	Amount           string
}
type ProposerHeaderRequest struct {
	Slot      uint64
	Proposer  string
	Timestamp uint64
}
type SlotPayloadRequest struct {
	Slot        uint64
	Proposer    string
	Transaction string
}
