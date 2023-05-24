package beaconclient

import (
	"github.com/attestantio/go-eth2-client/spec/capella"
	// capellaApi "github.com/attestantio/go-eth2-client/api/v1/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type PubkeyHex string

type ValidatorData struct {
	Index     uint64           `json:"index,string"`
	Balance   string           `json:"balance"` // In gwei.
	Status    string           `json:"status"`
	Validator ValidatorDetails `json:"validator"`
}

type ValidatorDetails struct {
	Pubkey string `json:"pubkey"`
}

type HeadEventData struct {
	Slot  uint64 `json:"slot,string"`
	Block string `json:"block"`
	State string `json:"state"`
}

type PayloadAttributesEventData struct {
	ProposerIndex     uint64            `json:"proposer_index,string"`
	ProposalSlot      uint64            `json:"proposal_slot,string"`
	ParentBlockNumber uint64            `json:"parent_block_number,string"`
	ParentBlockRoot   string            `json:"parent_block_root"`
	ParentBlockHash   string            `json:"parent_block_hash"`
	PayloadAttributes PayloadAttributes `json:"payload_attributes"`
}

type PayloadAttributes struct {
	Timestamp             uint64      `json:"timestamp,string"`
	PrevRandao            string      `json:"prev_randao"`
	SuggestedFeeRecipient string      `json:"suggested_fee_recipient"`
	Withdrawals           Withdrawals `json:"withdrawals"`
}

type GenesisData struct {
	GenesisTime           uint64 `json:"genesis_time,string"`
	GenesisValidatorsRoot string `json:"genesis_validators_root"`
	GenesisForkVersion    string `json:"genesis_fork_version"`
}

type SyncStatusData struct {
	HeadSlot  uint64 `json:"head_slot,string"`
	IsSyncing bool   `json:"is_syncing"`
}

type ProposerDutyData struct {
	PubkeyHex string `json:"pubkey"`
	Slot      uint64 `json:"slot,string"`
	Index     uint64 `json:"validator_index,string"`
}

type RandaoData struct {
	Randao common.Hash `json:"randao,string"`
}

type BeaconBlock capella.BeaconBlock

type BeaconBlockHeader struct {
	Slot          phase0.Slot           `json:"slot,string"`
	ProposerIndex phase0.ValidatorIndex `json:"proposer_index,string"`
	ParentRoot    string        `json:"parent_root" ssz-size:"32"`
	StateRoot     string           `json:"state_root" ssz-size:"32"`
	BodyRoot      string           `json:"body_root" ssz-size:"32"`
}

type SignedBeaconBlock capella.SignedBeaconBlock

type Withdrawal capella.Withdrawal

type Withdrawals []Withdrawal

type BlockHeaderData struct {
	Root string `json:"root"`
	Canonical bool `json:"canonical"`
	Header *SignedBeaconBlockHeader `json:"header"`
}

type SignedBeaconBlockHeader struct {
	Message *BeaconBlockHeader `json:"message"`
	Signature string `json:"signature"`
}
