package bundles

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type BuilderBundleEntry struct {
	ID         string    `db:"id"`
	InsertedAt time.Time `db:"inserted_at"`

	BundleHash        string `db:"bundle_hash"`
	Txs               string `db:"txs"`
	BlockNumber       uint64 `db:"block_number"`
	MinTimestamp      uint64 `db:"min_timestamp"`
	MaxTimestamp      uint64 `db:"max_timestamp"`
	RevertingTxHashes string `db:"reverting_tx_hashes"`

	BundleCreatorPubkey    string `db:"bundle_creator_pubkey"`
	BundleCreatorSignature string `db:"bundle_creator_signature"`

	BuilderPubkey    string `db:"builder_pubkey"`
	BuilderSignature string `db:"builder_signature"`

	BundleTransactionCount uint64 `db:"bundle_transaction_count"`
	BundleTotalGas         uint64 `db:"bundle_total_gas"`

	Added        bool   `db:"added"`
	Error        bool   `db:"error"`
	ErrorMessage string `db:"error_message"`
}

type BuilderBundle struct {
	ID         string
	BundleHash string

	Txs         []*types.Transaction
	BlockNumber uint64

	MinTimestamp uint64
	MaxTimestamp uint64

	RevertingTxHashes []*common.Hash

	BundleCreatorPubkey    string
	BundleCreatorSignature string

	BuilderPubkey    string
	BuilderSignature string

	BundleTransactionCount uint64
	BundleTotalGas         uint64

	Added        bool
	Error        bool
	ErrorMessage string

	Adding bool
}

func BuilderBundleToEntry(b *BuilderBundle) (*BuilderBundleEntry, error) {

	var txList []string
	for _, tx := range b.Txs {
		txBytes, err := tx.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("error marshalling tx: %v", err)
		}
		txList = append(txList, hexutil.Encode(txBytes))
	}

	var revertingTxHashes []string
	for _, txHash := range b.RevertingTxHashes {
		revertingTxHashes = append(revertingTxHashes, txHash.Hex())
	}

	return &BuilderBundleEntry{
		ID:                     b.ID,
		BundleHash:             b.BundleHash,
		Txs:                    strings.Join(txList, ","),
		BlockNumber:            b.BlockNumber,
		MinTimestamp:           b.MinTimestamp,
		MaxTimestamp:           b.MaxTimestamp,
		RevertingTxHashes:      strings.Join(revertingTxHashes, ","),
		BundleCreatorPubkey:    b.BundleCreatorPubkey,
		BundleCreatorSignature: b.BundleCreatorSignature,
		BuilderPubkey:          b.BuilderPubkey,
		BuilderSignature:       b.BuilderSignature,
		BundleTransactionCount: b.BundleTransactionCount,
		BundleTotalGas:         b.BundleTotalGas,
		Added:                  b.Added,
		Error:                  b.Error,
		ErrorMessage:           b.ErrorMessage,
	}, nil
}

func BuilderBundleEntryToBundle(b *BuilderBundleEntry) (*BuilderBundle, error) {

	var txs []*types.Transaction
	for _, txBytesEncoded := range strings.Split(b.Txs, ",") {
		txBytes, err := hexutil.Decode(txBytesEncoded)
		if err != nil {
			return nil, fmt.Errorf("error decoding tx bytes: %v", err)
		}
		var tx types.Transaction
		err = tx.UnmarshalBinary(txBytes)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling tx: %v", err)
		}
	}

	var revertingTxHashes []*common.Hash
	for _, txHash := range strings.Split(b.RevertingTxHashes, ",") {
		hash := common.HexToHash(txHash)
		revertingTxHashes = append(revertingTxHashes, &hash)
	}

	return &BuilderBundle{
		ID:                     b.ID,
		BundleHash:             b.BundleHash,
		Txs:                    txs,
		BlockNumber:            b.BlockNumber,
		MinTimestamp:           b.MinTimestamp,
		MaxTimestamp:           b.MaxTimestamp,
		RevertingTxHashes:      revertingTxHashes,
		BundleCreatorPubkey:    b.BundleCreatorPubkey,
		BundleCreatorSignature: b.BundleCreatorSignature,
		BuilderPubkey:          b.BuilderPubkey,
		BuilderSignature:       b.BuilderSignature,
		BundleTransactionCount: b.BundleTransactionCount,
		BundleTotalGas:         b.BundleTotalGas,
		Added:                  b.Added,
		Error:                  b.Error,
		ErrorMessage:           b.ErrorMessage,

		Adding: false,
	}, nil
}