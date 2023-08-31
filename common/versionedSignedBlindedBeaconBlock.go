package common

import (
	"errors"

	ssz "github.com/ferranbt/fastssz"

	bellatrix "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	capella "github.com/attestantio/go-eth2-client/api/v1/capella"
	deneb "github.com/attestantio/go-eth2-client/api/v1/deneb"
)

type VersionedSignedBlindedBeaconBlock struct {
	Bellatrix *bellatrix.SignedBlindedBeaconBlock `json:"bellatrix,omitempty"`
	Capella   *capella.SignedBlindedBeaconBlock   `json:"capella,omitempty"`
	Deneb     *deneb.SignedBlindedBeaconBlock     `json:"deneb,omitempty"`
}

func (v *VersionedSignedBlindedBeaconBlock) GetTree() (*ssz.Node, error) {
	if v.Deneb != nil {
		return v.Deneb.GetTree()
	}
	if v.Capella != nil {
		return v.Capella.GetTree()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.GetTree()
	}
	return nil, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) HashTreeRoot() ([32]byte, error) {
	if v.Deneb != nil {
		return v.Deneb.HashTreeRoot()
	}
	if v.Capella != nil {
		return v.Capella.HashTreeRoot()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.HashTreeRoot()
	}
	return [32]byte{}, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	if v.Deneb != nil {
		return v.Deneb.HashTreeRootWith(hh)
	}
	if v.Capella != nil {
		return v.Capella.HashTreeRootWith(hh)
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.HashTreeRootWith(hh)
	}
	return errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) MarshalJSON() ([]byte, error) {
	if v.Deneb != nil {
		return v.Deneb.MarshalJSON()
	}
	if v.Capella != nil {
		return v.Capella.MarshalJSON()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.MarshalJSON()
	}
	return nil, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) MarshalSSZ() ([]byte, error) {
	if v.Deneb != nil {
		return v.Deneb.MarshalSSZ()
	}
	if v.Capella != nil {
		return v.Capella.MarshalSSZ()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.MarshalSSZ()
	}
	return nil, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	if v.Deneb != nil {
		return v.Deneb.MarshalSSZTo(buf)
	}
	if v.Capella != nil {
		return v.Capella.MarshalSSZTo(buf)
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.MarshalSSZTo(buf)
	}
	return nil, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) MarshalYAML() ([]byte, error) {
	if v.Deneb != nil {
		return v.Deneb.MarshalYAML()
	}
	if v.Capella != nil {
		return v.Capella.MarshalYAML()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.MarshalYAML()
	}
	return nil, errors.New("no ExecutionPayload set")
}

func (v *VersionedSignedBlindedBeaconBlock) SizeSSZ() (size int) {
	if v.Deneb != nil {
		return v.Deneb.SizeSSZ()
	}
	if v.Capella != nil {
		return v.Capella.SizeSSZ()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.SizeSSZ()
	}
	return 0
}

func (v *VersionedSignedBlindedBeaconBlock) String() string {
	if v.Deneb != nil {
		return v.Deneb.String()
	}
	if v.Capella != nil {
		return v.Capella.String()
	}
	if v.Bellatrix != nil {
		return v.Bellatrix.String()
	}
	return "no ExecutionPayload set"
}

func (v *VersionedSignedBlindedBeaconBlock) UnmarshalJSON(input []byte) error {
	// The `UnmarshalJSON` function is a method of the `VersionedSignedBlindedBeaconBlock` struct. It is used to
	// deserialize JSON data into a `VersionedSignedBlindedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the JSON data.
	var err error

	v.Deneb = &deneb.SignedBlindedBeaconBlock{}
	err = v.Deneb.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBlindedBeaconBlock{}
	err = v.Capella.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBlindedBeaconBlock{}
	err = v.Bellatrix.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedSignedBlindedBeaconBlock) UnmarshalSSZ(buf []byte) error {
	// The `UnmarshalSSZ` function is a method of the `VersionedSignedBlindedBeaconBlock` struct. It is used to
	// deserialize SSZ data into a `VersionedSignedBlindedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the SSZ data.
	var err error

	v.Deneb = &deneb.SignedBlindedBeaconBlock{}
	err = v.Deneb.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBlindedBeaconBlock{}
	err = v.Capella.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBlindedBeaconBlock{}
	err = v.Bellatrix.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedSignedBlindedBeaconBlock) UnmarshalYAML(input []byte) error {
	// The `UnmarshalYAML` function is a method of the `VersionedSignedBlindedBeaconBlock` struct. It is used to
	// deserialize YAML data into a `VersionedSignedBlindedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the YAML data.
	var err error

	v.Deneb = &deneb.SignedBlindedBeaconBlock{}
	err = v.Deneb.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBlindedBeaconBlock{}
	err = v.Capella.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBlindedBeaconBlock{}
	err = v.Bellatrix.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}
