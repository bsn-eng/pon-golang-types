package relay

import (
	ssz "github.com/ferranbt/fastssz"
)

// HashTreeRoot ssz hashes the BidPayload object
func (b *BuilderBlockBid) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BidPayload object with a hasher
func (b *BuilderBlockBid) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (1) 'ProposerPubkey'
	hh.PutBytes(b.Pubkey[:])

	// Field (2) 'Value'
	valueBytes := b.Value.Bytes() // Big endian
	for i, j := 0, len(valueBytes)-1; i < j; i, j = i+1, j-1 {
		valueBytes[i], valueBytes[j] = valueBytes[j], valueBytes[i]
	} // Little endian
	hh.PutBytes(valueBytes)

	// Field (3) 'ExecutionPayloadHeader'
	headerRoot, err := b.ExecutionPayloadHeader.HashTreeRoot()
	if err != nil {
		return err
	}
	hh.PutBytes(headerRoot[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BidPayload object
func (b *BuilderBlockBid) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
