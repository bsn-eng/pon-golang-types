package builder

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	capella "github.com/attestantio/go-eth2-client/spec/capella"
	bundleTypes "github.com/bsn-eng/pon-golang-types/bundles"
	commonTypes "github.com/bsn-eng/pon-golang-types/common"
	rpbsTypes "github.com/bsn-eng/pon-golang-types/rpbs"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type BuilderPayloadAttributes struct {
	Timestamp             hexutil.Uint64              `json:"timestamp"`
	Random                gethCommon.Hash             `json:"prevRandao"`
	SuggestedFeeRecipient gethCommon.Address          `json:"suggestedFeeRecipient"`
	Slot                  uint64                      `json:"slot,string"`
	HeadHash              gethCommon.Hash             `json:"headHash"`
	BidAmount             *big.Int                    `json:"bidAmount"`
	GasLimit              uint64                      `json:"gasLimit,string"`
	Transactions          [][]byte                    `json:"transactions"`
	Withdrawals           types.Withdrawals           `json:"withdrawals"`
	NoMempoolTxs          bool                        `json:"noMempoolTxs,string"`
	PayoutPoolAddress     gethCommon.Address          `json:"payoutPoolAddress"`
	Bundles               []bundleTypes.BuilderBundle `json:"bundles"`
}

func (b *BuilderPayloadAttributes) UnmarshalJSON(data []byte) error {
	type Alias BuilderPayloadAttributes
	aux := &struct {
		BidAmount string `json:"bidAmount"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	b.BidAmount = new(big.Int)
	if _, ok := b.BidAmount.SetString(aux.BidAmount, 10); !ok {
		return fmt.Errorf("could not convert string to big.Int")
	}

	return nil
}

func (b *BuilderPayloadAttributes) MarshalJSON() ([]byte, error) {
	type Alias BuilderPayloadAttributes
	aux := &struct {
		BidAmount string `json:"bidAmount"`
		*Alias
	}{
		Alias:     (*Alias)(b),
		BidAmount: b.BidAmount.String(),
	}
	return json.Marshal(aux)
}

type PrivateTransactionsPayload struct {
	Transactions [][]byte `json:"transactions"`
}

type BuilderBlockBid struct {
	Signature      commonTypes.Signature      `json:"signature" ssz-size:"96"`
	Message        *BidPayload                `json:"message"`
	EcdsaSignature commonTypes.EcdsaSignature `json:"ecdsa_signature"`
}

type BlockBidResponse struct {
	RelayResponse      interface{}     `json:"relay_response"`
	BlockBid           BuilderBlockBid `json:"block_bid"`
	BidRequestTime     time.Time       `json:"bid_request_time"`
	BlockBuiltTime     time.Time       `json:"block_built_time"`
	BlockSubmittedTime time.Time       `json:"block_submitted_time"`
}

type BidPayload struct {
	Slot                 uint64                `json:"slot,string"`
	ParentHash           commonTypes.Hash      `json:"parent_hash" ssz-size:"32"`
	BlockHash            commonTypes.Hash      `json:"block_hash" ssz-size:"32"`
	BuilderPubkey        commonTypes.PublicKey `json:"builder_pubkey" ssz-size:"48"`
	ProposerPubkey       commonTypes.PublicKey `json:"proposer_pubkey" ssz-size:"48"`
	ProposerFeeRecipient commonTypes.Address   `json:"proposer_fee_recipient" ssz-size:"20"`
	GasLimit             uint64                `json:"gas_limit,string"`
	GasUsed              uint64                `json:"gas_used,string"`
	Value                *big.Int              `json:"value"`

	ExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"execution_payload_header"`
	Endpoint               string                          `json:"endpoint"`
	BuilderWalletAddress   commonTypes.Address             `json:"builder_wallet_address"`
	PayoutPoolTransaction  []byte                          `json:"payout_pool_transaction"`
	RPBS                   *rpbsTypes.EncodedRPBSSignature `json:"rpbs"`
	RPBSPubkey             string                          `json:"rpbs_pubkey"`
}

func (b *BidPayload) UnmarshalJSON(data []byte) error {
	type Alias BidPayload
	aux := &struct {
		Value string `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	b.Value = new(big.Int)
	if _, ok := b.Value.SetString(aux.Value, 10); !ok {
		return fmt.Errorf("could not convert string to big.Int")
	}

	return nil
}

func (b *BidPayload) MarshalJSON() ([]byte, error) {
	type Alias BidPayload
	aux := &struct {
		Value string `json:"value"`
		*Alias
	}{
		Alias: (*Alias)(b),
		Value: b.Value.String(),
	}
	return json.Marshal(aux)
}
