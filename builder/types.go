package builder

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
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
	NoMempoolTxs          bool                        `json:"noMempoolTxs,string"`
	PayoutPoolAddress     gethCommon.Address          `json:"payoutPoolAddress"`
	Withdrawals           types.Withdrawals           `json:"-"`
	Bundles               []bundleTypes.BuilderBundle `json:"-"`
}

// UnmarshalJSON implements the json.Unmarshaler interface purposefully for
// receiving a payload from the API.
func (b *BuilderPayloadAttributes) UnmarshalJSON(data []byte) error {
	type BuilderPayloadAttributesJSON struct {
		SuggestedFeeRecipient string   `json:"suggestedFeeRecipient"`
		Slot                  string   `json:"slot"`
		BidAmount             string   `json:"bidAmount"`
		Transactions          []string `json:"transactions"`
		NoMempoolTxs          string   `json:"noMempoolTxs"`
		PayoutPoolAddress     string   `json:"payoutPoolAddress"`
	}

	var aux BuilderPayloadAttributesJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	b.SuggestedFeeRecipient = gethCommon.Address{}
	if len(aux.SuggestedFeeRecipient) > 0 {
		if err := b.SuggestedFeeRecipient.UnmarshalText([]byte(aux.SuggestedFeeRecipient)); err != nil {
			return err
		}
	}

	b.Slot = 0
	if len(aux.Slot) > 0 {
		if _, err := fmt.Sscan(aux.Slot, &b.Slot); err != nil {
			return err
		}
	}

	b.BidAmount = big.NewInt(0)
	if len(aux.BidAmount) > 0 {
		if _, ok := b.BidAmount.SetString(aux.BidAmount, 10); !ok {
			return fmt.Errorf("failed to parse bid amount %s", aux.BidAmount)
		}
	}

	b.Transactions = make([][]byte, len(aux.Transactions))
	for i, tx := range aux.Transactions {
		b.Transactions[i] = []byte(tx)
	}

	b.NoMempoolTxs = false
	if len(aux.NoMempoolTxs) > 0 {
		if _, err := fmt.Sscan(aux.NoMempoolTxs, &b.NoMempoolTxs); err != nil {
			return err
		}
	}

	b.PayoutPoolAddress = gethCommon.Address{}
	if len(aux.PayoutPoolAddress) > 0 {
		if err := b.PayoutPoolAddress.UnmarshalText([]byte(aux.PayoutPoolAddress)); err != nil {
			return err
		}
	}

	return nil

}

type PrivateTransactionsPayload struct {
	Transactions [][]byte `json:"transactions"`
}

type BuilderBlockBid struct {
	Signature      commonTypes.Signature      `json:"signature" ssz-size:"96"`
	Message        *BidPayload                `json:"message"`
	EcdsaSignature commonTypes.EcdsaSignature `json:"ecdsa_signature"`
}

type SignedBuilderBlockBid struct {
	Message   *BuilderBlockBid `json:"message"`
	Signature phase0.BLSSignature `ssz-size:"96"`
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

	ExecutionPayloadHeader *commonTypes.VersionedExecutionPayloadHeader `json:"execution_payload_header"`
	Endpoint               string                          `json:"endpoint"`
	BuilderWalletAddress   commonTypes.Address             `json:"builder_wallet_address"`
	PayoutPoolTransaction  []byte                          `json:"payout_pool_transaction"`
	RPBS                   *rpbsTypes.EncodedRPBSSignature `json:"rpbs"`
	RPBSPubkey             string                          `json:"rpbs_pubkey"`
}
