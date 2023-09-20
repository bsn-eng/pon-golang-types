package relay

import (
	"math/big"
	ssz "github.com/ferranbt/fastssz"

	commonTypes "github.com/bsn-eng/pon-golang-types/common"

)

// MarshalSSZ ssz marshals the BuilderBid object
func (b *BuilderBlockBid) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BuilderBid object to a target array
func (b *BuilderBlockBid) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(84)

	// Offset (0) 'Header'
	dst = ssz.WriteOffset(dst, offset)
	if b.ExecutionPayloadHeader == nil {
		b.ExecutionPayloadHeader = new(commonTypes.VersionedExecutionPayloadHeader)
	}
	offset += b.ExecutionPayloadHeader.SizeSSZ()

	// Field (1) 'Value'
	value := b.Value.Bytes()
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	dst = append(dst, value...)

	// Field (2) 'Pubkey'
	dst = append(dst, b.Pubkey[:]...)

	// Field (0) 'Header'
	if dst, err = b.ExecutionPayloadHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BuilderBid object
func (b *BuilderBlockBid) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 84 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Header'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 84 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Value'
	value := buf[4:36]
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	b.Value = new(big.Int)
	b.Value.SetBytes(value)

	// Field (2) 'Pubkey'
	copy(b.Pubkey[:], buf[36:84])

	// Field (0) 'Header'
	{
		buf = tail[o0:]
		if b.ExecutionPayloadHeader == nil {
			b.ExecutionPayloadHeader = new(commonTypes.VersionedExecutionPayloadHeader)
		}
		if err = b.ExecutionPayloadHeader.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BuilderBid object
func (b *BuilderBlockBid) SizeSSZ() (size int) {
	size = 84

	// Field (0) 'Header'
	if b.ExecutionPayloadHeader == nil {
		b.ExecutionPayloadHeader = new(commonTypes.VersionedExecutionPayloadHeader)
	}
	size += b.ExecutionPayloadHeader.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BuilderBid object
func (b *BuilderBlockBid) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BuilderBid object with a hasher
func (b *BuilderBlockBid) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Header'
	if err = b.ExecutionPayloadHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Value'
	value := b.Value.Bytes()
	for i, j := 0, len(value)-1; i < j; i, j = i+1, j-1 {
		value[i], value[j] = value[j], value[i]
	}
	hh.PutBytes(value)

	// Field (2) 'Pubkey'
	hh.PutBytes(b.Pubkey[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BuilderBid object
func (b *BuilderBlockBid) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}