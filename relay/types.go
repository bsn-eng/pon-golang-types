package relay

import (
	"github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type Address [20]byte
type Signature phase0.BLSSignature
type EcdsaAddress [20]byte
type EcdsaSignature [65]byte
type Hash [32]byte
type PublicKey [48]byte
type Transaction []byte

type CapellaBuilderSubmitBlockRequest struct {
	Signature      Signature      `json:"signature" ssz-size:"96"`
	Message        *BidTrace      `json:"message"`
	EcdsaSignature EcdsaSignature `json:"ecdsa_signature"`
}

type BidTrace struct {
	Slot                 uint64    `json:"slot,string"`
	ParentHash           Hash      `json:"parent_hash" ssz-size:"32"`
	BlockHash            Hash      `json:"block_hash" ssz-size:"32"`
	BuilderPubkey        PublicKey `json:"builder_pubkey" ssz-size:"48"`
	ProposerPubkey       PublicKey `json:"proposer_pubkey" ssz-size:"48"`
	ProposerFeeRecipient Address   `json:"proposer_fee_recipient" ssz-size:"20"`
	GasLimit             uint64    `json:"gas_limit,string"`
	GasUsed              uint64    `json:"gas_used,string"`
	Value                uint64    `json:"value" ssz-size:"32"`

	ExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"execution_payload_header"`
	Endpoint               string                          `json:"endpoint"`
	BuilderWalletAddress   Address                         `json:"builder_wallet_address"`
	PayoutPoolTransaction  []byte                          `json:"payout_pool_transaction"`
	RPBS                   *EncodedRPBSSignature           `json:"rpbs"`
	RPBSPubkey             string                          `json:"rpbs_pubkey"`
}

type EncodedRPBSSignature struct {
	Z1Hat string `json:"z1_hat"`
	C1Hat string `json:"c1_hat"`
	S1Hat string `json:"s1_hat"`
	C2Hat string `json:"c2_hat"`
	S2Hat string `json:"s2_hat"`
	M1Hat string `json:"m1_hat"`
}
