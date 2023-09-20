package relay

import (
	"encoding/json"
	"sync"
	"math/big"
	"errors"

	commonTypes "github.com/bsn-eng/pon-golang-types/common"

	"github.com/attestantio/go-eth2-client/spec/phase0"
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

func (h BLSPubKey) String() string {
	return hexutil.Bytes(h[:]).String()
}

func (e EcdsaAddress) String() string {
	return hexutil.Bytes(e[:]).String()
}

func (e EcdsaSignature) String() string {
	return hexutil.Bytes(e[:]).String()
}

func (e Address) String() string {
	return hexutil.Bytes(e[:]).String()
}

func (t Transaction) String() string {
	transaction, _ := json.Marshal(t)
	return string(transaction)
}

type ValidatorIndexes struct {
	Mu                   sync.Mutex
	ValidatorPubkeyIndex map[string]uint64
	ValidatorIndexPubkey map[uint64]string
}

// SignedBuilderBlockBid is a signed BuilderBlockBid similar to builder.SignedBuilderBlockBid
type SignedBuilderBlockBid struct {
	Message   *BuilderBlockBid    `json:"message"`
	Signature phase0.BLSSignature `ssz-size:"96"`
}

// BuilderBlockBid is a BuilderBlockBid similar to builder.BuilderBlockBid
// This is just leaner with only necessary fields passed to valiator proxy software
type BuilderBlockBid struct {
	Pubkey phase0.BLSPubKey `json:"pubkey" ssz-size:"48"`
	// json feild name has been changed from proposer_pubkey to pubkey for mevBoost

	Value *big.Int `json:"value"`

	ExecutionPayloadHeader *commonTypes.VersionedExecutionPayloadHeader `json:"header"`
	// json feild name has been changed from execution_payload_header to header for mevBoost
}

type builderBlockBidJSON struct {
	Pubkey string `json:"pubkey" ssz-size:"48"`
	Value          string `json:"value"`
	ExecutionPayloadHeader *commonTypes.VersionedExecutionPayloadHeader `json:"header"`
}

func (b *BuilderBlockBid) MarshalJSON() ([]byte, error) {
	return json.Marshal(&builderBlockBidJSON{
		Pubkey: b.Pubkey.String(),
		Value:          b.Value.String(),
		ExecutionPayloadHeader: b.ExecutionPayloadHeader,
	})
}

func (b *BuilderBlockBid) UnmarshalJSON(input []byte) error {
	var data builderBlockBidJSON
	if err := json.Unmarshal(input, &data); err != nil {
		return err
	}
	return b.unpack(&data)
}

func (b *BuilderBlockBid) unpack(data *builderBlockBidJSON) error {
	if data.Pubkey == "" {
		return errors.New("pubkey missing")
	}
	pubkey, err := hexutil.Decode(data.Pubkey)
	if err != nil {
		return err
	}
	copy(b.Pubkey[:], pubkey)

	if data.Value == "" {
		return errors.New("value missing")
	}
	value, ok := big.NewInt(0).SetString(data.Value, 10)
	if !ok {
		return errors.New("invalid value for value")
	}
	b.Value = value

	if data.ExecutionPayloadHeader == nil {
		return errors.New("header missing")
	}
	b.ExecutionPayloadHeader = data.ExecutionPayloadHeader

	return nil

}
