package common

import (
	"errors"

	ssz "github.com/ferranbt/fastssz"

	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
)

type VersionedSignedBeaconBlock struct {
	Bellatrix *bellatrix.SignedBeaconBlock `json:"bellatrix,omitempty"`
	Capella   *capella.SignedBeaconBlock   `json:"capella,omitempty"`
	Deneb     *deneb.SignedBeaconBlock     `json:"deneb,omitempty"`
}

func (v *VersionedSignedBeaconBlock) GetTree() (*ssz.Node, error) {
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

func (v *VersionedSignedBeaconBlock) HashTreeRoot() ([32]byte, error) {
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

func (v *VersionedSignedBeaconBlock) HashTreeRootWith(hh ssz.HashWalker) (err error) {
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

func (v *VersionedSignedBeaconBlock) MarshalJSON() ([]byte, error) {
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

func (v *VersionedSignedBeaconBlock) MarshalSSZ() ([]byte, error) {
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

func (v *VersionedSignedBeaconBlock) MarshalSSZTo(buf []byte) (dst []byte, err error) {
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

func (v *VersionedSignedBeaconBlock) MarshalYAML() ([]byte, error) {
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

func (v *VersionedSignedBeaconBlock) SizeSSZ() (size int) {
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

func (v *VersionedSignedBeaconBlock) String() string {
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

func (v *VersionedSignedBeaconBlock) UnmarshalJSON(input []byte) error {
	// The `UnmarshalJSON` function is a method of the `VersionedSignedBeaconBlock` struct. It is used to
	// deserialize JSON data into a `VersionedSignedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the JSON data.
	var err error

	v.Deneb = &deneb.SignedBeaconBlock{}
	err = v.Deneb.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBeaconBlock{}
	err = v.Capella.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBeaconBlock{}
	err = v.Bellatrix.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedSignedBeaconBlock) UnmarshalSSZ(buf []byte) error {
	// The `UnmarshalSSZ` function is a method of the `VersionedSignedBeaconBlock` struct. It is used to
	// deserialize SSZ data into a `VersionedSignedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the SSZ data.
	var err error

	v.Deneb = &deneb.SignedBeaconBlock{}
	err = v.Deneb.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBeaconBlock{}
	err = v.Capella.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBeaconBlock{}
	err = v.Bellatrix.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedSignedBeaconBlock) UnmarshalYAML(input []byte) error {
	// The `UnmarshalYAML` function is a method of the `VersionedSignedBeaconBlock` struct. It is used to
	// deserialize YAML data into a `VersionedSignedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the YAML data.
	var err error

	v.Deneb = &deneb.SignedBeaconBlock{}
	err = v.Deneb.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.SignedBeaconBlock{}
	err = v.Capella.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.SignedBeaconBlock{}
	err = v.Bellatrix.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}
