package beaconclient

import (

	"sync"
	
	"github.com/ethereum/go-ethereum/common"
)

type BeaconClient struct {
	endpoint      string
	slotsInEpoch  uint64
	secondsInSlot uint64
	currentSlot   uint64
	currentEpoch  uint64

	currentHead HeadEventData
	randaoMap   map[uint64]common.Hash

	mu              sync.Mutex
	slotProposerMap map[uint64]ValidatorData
	slotPayloadAttributesMap  map[uint64]PayloadAttributesEventData

	closeCh chan struct{}
}