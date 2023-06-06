package common

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var ErrLength = fmt.Errorf("incorrect byte length")

/*
Signature types and methods
*/
type Signature [96]byte

func (s Signature) MarshalText() ([]byte, error) {
	return hexutil.Bytes(s[:]).MarshalText()
}

func (s *Signature) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalJSON(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s *Signature) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalText(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s Signature) String() string {
	return hexutil.Bytes(s[:]).String()
}

func (s *Signature) FromSlice(x []byte) error {
	if len(x) != 96 {
		return ErrLength
	}
	copy(s[:], x)
	return nil
}

/*
ECDSA Signature types and methods
*/
type EcdsaSignature [65]byte

func (s EcdsaSignature) MarshalText() ([]byte, error) {
	return hexutil.Bytes(s[:]).MarshalText()
}

func (s *EcdsaSignature) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalJSON(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s *EcdsaSignature) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(s[:])
	err := b.UnmarshalText(input)
	if err != nil {
		return err
	}
	return s.FromSlice(b)
}

func (s EcdsaSignature) String() string {
	return hexutil.Bytes(s[:]).String()
}

func (s *EcdsaSignature) FromSlice(x []byte) error {
	if len(x) != 65 {
		return ErrLength
	}
	copy(s[:], x)
	return nil
}

/*
Hash types and methods
*/
type (
	Hash [32]byte
	Root = Hash
)

func (h Hash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

func (h *Hash) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(h[:])
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return h.FromSlice(b)
}

func (h *Hash) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(h[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
	return h.FromSlice(b)
}

func (h *Hash) FromSlice(x []byte) error {
	if len(x) != 32 {
		return ErrLength
	}
	copy(h[:], x)
	return nil
}

func (h Hash) String() string {
	return hexutil.Bytes(h[:]).String()
}

/*
Address types and methods
*/
type Address [20]byte

func (a Address) MarshalText() ([]byte, error) {
	return hexutil.Bytes(a[:]).MarshalText()
}

func (a *Address) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(a[:])
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return a.FromSlice(b)
}

func (a *Address) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(a[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
	return a.FromSlice(b)
}

func (a Address) String() string {
	return hexutil.Bytes(a[:]).String()
}

func (a *Address) FromSlice(x []byte) error {
	if len(x) != 20 {
		return ErrLength
	}
	copy(a[:], x)
	return nil
}

/*
Public key types and methods
*/
type PublicKey [48]byte

func (p PublicKey) MarshalText() ([]byte, error) {
	return hexutil.Bytes(p[:]).MarshalText()
}

func (p *PublicKey) UnmarshalJSON(input []byte) error {
	b := hexutil.Bytes(p[:])
	if err := b.UnmarshalJSON(input); err != nil {
		return err
	}
	return p.FromSlice(b)
}

func (p *PublicKey) UnmarshalText(input []byte) error {
	b := hexutil.Bytes(p[:])
	if err := b.UnmarshalText(input); err != nil {
		return err
	}
	return p.FromSlice(b)
}

func (p PublicKey) String() string {
	return hexutil.Bytes(p[:]).String()
}

func (p *PublicKey) FromSlice(x []byte) error {
	if len(x) != 48 {
		return ErrLength
	}
	copy(p[:], x)
	return nil
}

func HexToPubkey(s string) (ret PublicKey, err error) {
	err = ret.UnmarshalText([]byte(s))
	return ret, err
}
