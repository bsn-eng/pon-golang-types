package rpbs

import (
	ssz "github.com/ferranbt/fastssz"
)

type RPBSCommitMessage struct {
	BuilderWalletAddress string `json:"builderWalletAddress"`
	Slot                 uint64 `json:"slot"`
	Amount               uint64 `json:"amount"`
	PayoutTxBytes        string `json:"payoutTxBytes"`
	TxBytes              []string `json:"txBytes"`
}

type EncodedRPBSSignature struct {
	Z1Hat string `json:"z1Hat"`
	C1Hat string `json:"c1Hat"`
	S1Hat string `json:"s1Hat"`
	C2Hat string `json:"c2Hat"`
	S2Hat string `json:"s2Hat"`
	M1Hat string `json:"m1Hat"`
}

// HashTreeRoot ssz hashes the EncodedRPBSSignature object
func (e *EncodedRPBSSignature) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(e)
}

// HashTreeRootWith ssz hashes the EncodedRPBSSignature object with a hasher
func (e *EncodedRPBSSignature) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Z1Hat'
	hh.PutBytes([]byte(e.Z1Hat))

	// Field (1) 'C1Hat'
	hh.PutBytes([]byte(e.C1Hat))

	// Field (2) 'S1Hat'
	hh.PutBytes([]byte(e.S1Hat))

	// Field (3) 'C2Hat'
	hh.PutBytes([]byte(e.C2Hat))

	// Field (4) 'S2Hat'
	hh.PutBytes([]byte(e.S2Hat))

	// Field (5) 'M1Hat'
	hh.PutBytes([]byte(e.M1Hat))

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the EncodedRPBSSignature object
func (e *EncodedRPBSSignature) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(e)
}