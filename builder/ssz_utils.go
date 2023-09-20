package builder

import (
	ssz "github.com/ferranbt/fastssz"
)

// HashTreeRoot ssz hashes the BidPayload object
func (b *BidPayload) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BidPayload object with a hasher
func (b *BidPayload) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (1) 'ParentHash'
	hh.PutBytes(b.ParentHash[:])

	// Field (2) 'BlockHash'
	hh.PutBytes(b.BlockHash[:])

	// Field (3) 'BuilderPubkey'
	hh.PutBytes(b.BuilderPubkey[:])

	// Field (4) 'ProposerPubkey'
	hh.PutBytes(b.ProposerPubkey[:])

	// Field (5) 'ProposerFeeRecipient'
	hh.PutBytes(b.ProposerFeeRecipient[:])

	// Field (6) 'GasLimit'
	hh.PutUint64(b.GasLimit)

	// Field (7) 'GasUsed'
	hh.PutUint64(b.GasUsed)

	// Field (8) 'Value'
	valueBytes := b.Value.Bytes() // Big endian
	for i, j := 0, len(valueBytes)-1; i < j; i, j = i+1, j-1 {
		valueBytes[i], valueBytes[j] = valueBytes[j], valueBytes[i]
	} // Little endian
	hh.PutBytes(valueBytes)

	// Field (9) 'ExecutionPayloadHeader'
	headerRoot, err := b.ExecutionPayloadHeader.HashTreeRoot()
	if err != nil {
		return err
	}
	hh.PutBytes(headerRoot[:])

	// Field (10) 'Endpoint'
	hh.PutBytes([]byte(b.Endpoint))

	// Field (11) 'BuilderWalletAddress'
	hh.PutBytes(b.BuilderWalletAddress[:])

	// Field (12) 'PayoutPoolTransaction'
	hh.PutBytes(b.PayoutPoolTransaction[:])

	// Field (13) 'RPBS'
	rpbsHash, err := b.RPBS.HashTreeRoot()
	if err != nil {
		return err
	}
	hh.PutBytes(rpbsHash[:])

	// Field (14) 'RPBSPubkey'
	hh.PutBytes([]byte(b.RPBSPubkey))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BidPayload object
func (b *BidPayload) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}