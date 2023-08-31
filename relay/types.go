package relay

import (
	"encoding/json"
	"sync"

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
