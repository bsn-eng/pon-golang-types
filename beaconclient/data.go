package beaconclient

import (
	"github.com/attestantio/go-eth2-client/spec/capella"
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
	Timestamp             uint64                `json:"timestamp,string"`
	PrevRandao            string                `json:"prev_randao"`
	SuggestedFeeRecipient string                `json:"suggested_fee_recipient"`
	Withdrawals           []*capella.Withdrawal `json:"withdrawals"`
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
	Slot      uint64 `json:"slot"`
	Index     uint64 `json:"validator_index"`
}

type WithdrawalData capella.Withdrawal

type RandaoData struct {
	Randao common.Hash `json:"randao"`
}

type BeaconBlock capella.BeaconBlock

type SignedBeaconBlock capella.SignedBeaconBlock
