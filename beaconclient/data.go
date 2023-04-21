package beaconclient

import (

	"github.com/attestantio/go-eth2-client/spec/capella"
)

type PubkeyHex string

type ValidatorData struct {
	Pubkey         PubkeyHex
	ValidatorIndex uint64
}

type GetGenesisResponse struct {
	Data struct {
		GenesisTime           uint64 `json:"genesis_time,string"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		GenesisForkVersion    string `json:"genesis_fork_version"`
	}
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
	Timestamp             uint64                `json:"timestamp,string"`
	PrevRandao            string                `json:"prev_randao"`
	SuggestedFeeRecipient string                `json:"suggested_fee_recipient"`
	Withdrawals           []*capella.Withdrawal `json:"withdrawals"`
}