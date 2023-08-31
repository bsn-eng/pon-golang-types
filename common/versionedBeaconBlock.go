package common

import (
	"errors"

	ssz "github.com/ferranbt/fastssz"

	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
)

type VersionedBeaconBlock struct {
	Bellatrix *bellatrix.BeaconBlock `json:"bellatrix,omitempty"`
	Capella   *capella.BeaconBlock   `json:"capella,omitempty"`
	Deneb     *deneb.BeaconBlock     `json:"deneb,omitempty"`
}

func (v *VersionedBeaconBlock) GetTree() (*ssz.Node, error) {
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

func (v *VersionedBeaconBlock) HashTreeRoot() ([32]byte, error) {
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

func (v *VersionedBeaconBlock) HashTreeRootWith(hh ssz.HashWalker) (err error) {
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

func (v *VersionedBeaconBlock) MarshalJSON() ([]byte, error) {
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

func (v *VersionedBeaconBlock) MarshalSSZ() ([]byte, error) {
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

func (v *VersionedBeaconBlock) MarshalSSZTo(buf []byte) (dst []byte, err error) {
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

func (v *VersionedBeaconBlock) MarshalYAML() ([]byte, error) {
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

func (v *VersionedBeaconBlock) SizeSSZ() (size int) {
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

func (v *VersionedBeaconBlock) String() string {
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

func (v *VersionedBeaconBlock) UnmarshalJSON(input []byte) error {
	// The `UnmarshalJSON` function is a method of the `VersionedBeaconBlock` struct. It is used to
	// deserialize JSON data into a `VersionedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the JSON data.
	var err error

	v.Deneb = &deneb.BeaconBlock{}
	err = v.Deneb.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.BeaconBlock{}
	err = v.Capella.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.BeaconBlock{}
	err = v.Bellatrix.UnmarshalJSON(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedBeaconBlock) UnmarshalSSZ(buf []byte) error {
	// The `UnmarshalSSZ` function is a method of the `VersionedBeaconBlock` struct. It is used to
	// deserialize SSZ data into a `VersionedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the SSZ data.
	var err error

	v.Deneb = &deneb.BeaconBlock{}
	err = v.Deneb.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.BeaconBlock{}
	err = v.Capella.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.BeaconBlock{}
	err = v.Bellatrix.UnmarshalSSZ(buf)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}

func (v *VersionedBeaconBlock) UnmarshalYAML(input []byte) error {
	// The `UnmarshalYAML` function is a method of the `VersionedBeaconBlock` struct. It is used to
	// deserialize YAML data into a `VersionedBeaconBlock` object.

	// Type is forkversion naive so we need to try each type in reverse
	// fork version order.  This is because the fork version is not
	// included in the YAML data.
	var err error

	v.Deneb = &deneb.BeaconBlock{}
	err = v.Deneb.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Deneb = nil

	v.Capella = &capella.BeaconBlock{}
	err = v.Capella.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Capella = nil

	v.Bellatrix = &bellatrix.BeaconBlock{}
	err = v.Bellatrix.UnmarshalYAML(input)
	if err == nil {
		return nil
	}
	v.Bellatrix = nil

	return errors.New("unsupported ExecutionPayload type")

}
