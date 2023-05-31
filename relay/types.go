package relay

import (
	"encoding/json"

	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	rpbsTypes "github.com/bsn-eng/pon-golang-types/rpbs"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address [20]byte
type Signature phase0.BLSSignature
type EcdsaAddress [20]byte
type EcdsaSignature [65]byte
type Hash [32]byte
type BLSPubKey [48]byte
type Transaction []byte

func (h Hash) String() string {
	return hexutil.Bytes(h[:]).String()
}

func (e EcdsaAddress) String() string {
	return hexutil.Bytes(e[:]).String()
}

func (e EcdsaSignature) String() string {
	return hexutil.Bytes(e[:]).String()
}

func (t Transaction) String() string {
	transaction, _ := json.Marshal(t)
	return string(transaction)
}

type BuilderSubmitBlockRequest struct {
	Signature      Signature      `json:"signature" ssz-size:"96"`
	Message        *BidTrace      `json:"message"`
	EcdsaSignature EcdsaSignature `json:"ecdsa_signature"`
}

type BidTrace struct {
	Slot                 uint64    `json:"slot,string"`
	ParentHash           Hash      `json:"parent_hash" ssz-size:"32"`
	BlockHash            Hash      `json:"block_hash" ssz-size:"32"`
	BuilderPubkey        BLSPubKey `json:"builder_pubkey" ssz-size:"48"`
	ProposerPubkey       BLSPubKey `json:"proposer_pubkey" ssz-size:"48"`
	ProposerFeeRecipient Address   `json:"proposer_fee_recipient" ssz-size:"20"`
	GasLimit             uint64    `json:"gas_limit,string"`
	GasUsed              uint64    `json:"gas_used,string"`
	Value                uint64    `json:"value" ssz-size:"32"`

	ExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"execution_payload_header"`
	Endpoint               string                          `json:"endpoint"`
	BuilderWalletAddress   Address                         `json:"builder_wallet_address"`
	PayoutPoolTransaction  []byte                          `json:"payout_pool_transaction"`
	RPBS                   *rpbsTypes.EncodedRPBSSignature `json:"rpbs"`
	RPBSPubkey             string                          `json:"rpbs_pubkey"`
}
