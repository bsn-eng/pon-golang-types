package bundles

import (
	ssz "github.com/ferranbt/fastssz"
)

// HashTreeRoot ssz hashes the main contents of the BuilderBundle object
func (b *BuilderBundle) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the main contents of the main contents of the BuilderBundle with a hasher
func (b *BuilderBundle) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Txs'
	for i := 0; i < len(b.Txs); i++ {
		bytes, err := b.Txs[i].MarshalBinary()
		if err != nil {
			return err
		}
		hh.PutBytes(bytes)
	}

	// Field (1) 'BlockNumber'
	hh.PutUint64(b.BlockNumber)

	// Field (2) 'MinTimestamp'
	hh.PutUint64(b.MinTimestamp)

	// Field (3) 'MaxTimestamp'
	hh.PutUint64(b.MaxTimestamp)

	// Field (4) 'RevertingTxHashes'
	for i := 0; i < len(b.RevertingTxHashes); i++ {
		hh.PutBytes(b.RevertingTxHashes[i][:])
	}

	// Field (5) 'BundleTransactionCount'
	hh.PutUint64(b.BundleTransactionCount)

	// Field (6) 'BundleTotalGas'
	hh.PutBytes(b.BundleTotalGas.Bytes())

	// Field (7) 'BundleDateTime'
	hh.PutUint64(uint64(b.BundleDateTime.Unix()))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the main contents of the BuilderBundle object
func (b *BuilderBundle) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}