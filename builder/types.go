package builder

import (
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
	BidAmount             uint64                      `json:"bidAmount,string"`
	GasLimit              uint64                      `json:"gasLimit,string"`
	Transactions          [][]byte                    `json:"transactions"`
	Withdrawals           types.Withdrawals           `json:"withdrawals"`
	NoMempoolTxs          bool                        `json:"noMempoolTxs,string"`
	PayoutPoolAddress     gethCommon.Address          `json:"payoutPoolAddress"`
	Bundles               []bundleTypes.BuilderBundle `json:"bundles"`
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
	RelayResponse interface{}     `json:"relay_response"`
	Error         error           `json:"error"`
	BlockBid      BuilderBlockBid `json:"block_bid"`
	BidRequestTime time.Time       `json:"bid_request_time"`
	BlockBuiltTime time.Time       `json:"block_built_time"`
	RelaySubmissionTime time.Time  `json:"relay_submission_time"`
}

type BuilderBidRelay struct {
	BidID             string `json:"bid_id"`
	HighestBidValue   uint64 `json:"highest_bid_value,string"`
	HighestBidBuilder string `json:"highest_bid_builder"`
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
	Value                uint64                `json:"value,string"`

	ExecutionPayloadHeader *capella.ExecutionPayloadHeader `json:"execution_payload_header"`
	Endpoint               string                          `json:"endpoint"`
	BuilderWalletAddress   commonTypes.Address             `json:"builder_wallet_address"`
	PayoutPoolTransaction  []byte                          `json:"payout_pool_transaction"`
	RPBS                   *rpbsTypes.EncodedRPBSSignature `json:"rpbs"`
	RPBSPubkey             string                          `json:"rpbs_pubkey"`
}
