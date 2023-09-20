package common

import (
	"errors"
	"math/big"

	"github.com/attestantio/go-eth2-client/spec"

	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseSignedBeaconBlock struct {
	Message   *BaseBeaconBlock
	Signature phase0.BLSSignature `ssz-size:"96"`
}

func ConstructSignedBeaconBlock(
	forkVersion string,
	signedBeaconBlock BaseSignedBeaconBlock,
) (VersionedSignedBeaconBlock, error) {

	res := VersionedSignedBeaconBlock{}

	VersionedBeaconBlock, err := ConstructBeaconBlock(forkVersion, *signedBeaconBlock.Message)
	if err != nil {
		return res, err
	}

	switch forkVersion {
	case spec.DataVersionBellatrix.String():
		res.Bellatrix = &bellatrix.SignedBeaconBlock{
			Message:   VersionedBeaconBlock.Bellatrix,
			Signature: signedBeaconBlock.Signature,
		}
	case spec.DataVersionCapella.String():
		res.Capella = &capella.SignedBeaconBlock{
			Message:   VersionedBeaconBlock.Capella,
			Signature: signedBeaconBlock.Signature,
		}
	case spec.DataVersionDeneb.String():
		res.Deneb = &deneb.SignedBeaconBlock{
			Message:   VersionedBeaconBlock.Deneb,
			Signature: signedBeaconBlock.Signature,
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (b *VersionedSignedBeaconBlock) ToBaseSignedBeaconBlock() (BaseSignedBeaconBlock, error) {
	res := BaseSignedBeaconBlock{}

	switch {
	case b.Bellatrix != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Bellatrix.Message.Body.ExecutionPayload.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Bellatrix.Message.Body.ExecutionPayload.BaseFeePerGas[len(b.Bellatrix.Message.Body.ExecutionPayload.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}
		res.Message = &BaseBeaconBlock{
			Slot:          b.Bellatrix.Message.Slot,
			ProposerIndex: b.Bellatrix.Message.ProposerIndex,
			ParentRoot:    b.Bellatrix.Message.ParentRoot,
			StateRoot:     b.Bellatrix.Message.StateRoot,
			Body: &BaseBeaconBlockBody{
				RANDAOReveal:      b.Bellatrix.Message.Body.RANDAOReveal,
				ETH1Data:          b.Bellatrix.Message.Body.ETH1Data,
				Graffiti:          b.Bellatrix.Message.Body.Graffiti,
				ProposerSlashings: b.Bellatrix.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Bellatrix.Message.Body.AttesterSlashings,
				Attestations:      b.Bellatrix.Message.Body.Attestations,
				Deposits:          b.Bellatrix.Message.Body.Deposits,
				VoluntaryExits:    b.Bellatrix.Message.Body.VoluntaryExits,
				SyncAggregate:     b.Bellatrix.Message.Body.SyncAggregate,
				ExecutionPayload: &BaseExecutionPayload{
					ParentHash:    b.Bellatrix.Message.Body.ExecutionPayload.ParentHash,
					FeeRecipient:  b.Bellatrix.Message.Body.ExecutionPayload.FeeRecipient,
					StateRoot:     b.Bellatrix.Message.Body.ExecutionPayload.StateRoot,
					ReceiptsRoot:  b.Bellatrix.Message.Body.ExecutionPayload.ReceiptsRoot,
					LogsBloom:     b.Bellatrix.Message.Body.ExecutionPayload.LogsBloom,
					PrevRandao:    b.Bellatrix.Message.Body.ExecutionPayload.PrevRandao,
					BlockNumber:   b.Bellatrix.Message.Body.ExecutionPayload.BlockNumber,
					GasLimit:      b.Bellatrix.Message.Body.ExecutionPayload.GasLimit,
					GasUsed:       b.Bellatrix.Message.Body.ExecutionPayload.GasUsed,
					Timestamp:     b.Bellatrix.Message.Body.ExecutionPayload.Timestamp,
					ExtraData:     b.Bellatrix.Message.Body.ExecutionPayload.ExtraData,
					BaseFeePerGas: baseFeePerGas,
					BlockHash:     b.Bellatrix.Message.Body.ExecutionPayload.BlockHash,
					Transactions:  b.Bellatrix.Message.Body.ExecutionPayload.Transactions,
				},
			},
		}
		res.Signature = b.Bellatrix.Signature
	case b.Capella != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Capella.Message.Body.ExecutionPayload.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Capella.Message.Body.ExecutionPayload.BaseFeePerGas[len(b.Capella.Message.Body.ExecutionPayload.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.Message = &BaseBeaconBlock{
			Slot:          b.Capella.Message.Slot,
			ProposerIndex: b.Capella.Message.ProposerIndex,
			ParentRoot:    b.Capella.Message.ParentRoot,
			StateRoot:     b.Capella.Message.StateRoot,
			Body: &BaseBeaconBlockBody{
				RANDAOReveal:      b.Capella.Message.Body.RANDAOReveal,
				ETH1Data:          b.Capella.Message.Body.ETH1Data,
				Graffiti:          b.Capella.Message.Body.Graffiti,
				ProposerSlashings: b.Capella.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Capella.Message.Body.AttesterSlashings,
				Attestations:      b.Capella.Message.Body.Attestations,
				Deposits:          b.Capella.Message.Body.Deposits,
				VoluntaryExits:    b.Capella.Message.Body.VoluntaryExits,
				SyncAggregate:     b.Capella.Message.Body.SyncAggregate,
				ExecutionPayload: &BaseExecutionPayload{
					ParentHash:    b.Capella.Message.Body.ExecutionPayload.ParentHash,
					FeeRecipient:  b.Capella.Message.Body.ExecutionPayload.FeeRecipient,
					StateRoot:     b.Capella.Message.Body.ExecutionPayload.StateRoot,
					ReceiptsRoot:  b.Capella.Message.Body.ExecutionPayload.ReceiptsRoot,
					LogsBloom:     b.Capella.Message.Body.ExecutionPayload.LogsBloom,
					PrevRandao:    b.Capella.Message.Body.ExecutionPayload.PrevRandao,
					BlockNumber:   b.Capella.Message.Body.ExecutionPayload.BlockNumber,
					GasLimit:      b.Capella.Message.Body.ExecutionPayload.GasLimit,
					GasUsed:       b.Capella.Message.Body.ExecutionPayload.GasUsed,
					Timestamp:     b.Capella.Message.Body.ExecutionPayload.Timestamp,
					ExtraData:     b.Capella.Message.Body.ExecutionPayload.ExtraData,
					BaseFeePerGas: baseFeePerGas,
					BlockHash:     b.Capella.Message.Body.ExecutionPayload.BlockHash,
					Transactions:  b.Capella.Message.Body.ExecutionPayload.Transactions,
					Withdrawals:   b.Capella.Message.Body.ExecutionPayload.Withdrawals,
				},
				BLSToExecutionChanges: b.Capella.Message.Body.BLSToExecutionChanges,
			},
		}
		res.Signature = b.Capella.Signature
	case b.Deneb != nil:

		res.Message = &BaseBeaconBlock{
			Slot:          b.Deneb.Message.Slot,
			ProposerIndex: b.Deneb.Message.ProposerIndex,
			ParentRoot:    b.Deneb.Message.ParentRoot,
			StateRoot:     b.Deneb.Message.StateRoot,
			Body: &BaseBeaconBlockBody{
				RANDAOReveal:      b.Deneb.Message.Body.RANDAOReveal,
				ETH1Data:          b.Deneb.Message.Body.ETH1Data,
				Graffiti:          b.Deneb.Message.Body.Graffiti,
				ProposerSlashings: b.Deneb.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Deneb.Message.Body.AttesterSlashings,
				Attestations:      b.Deneb.Message.Body.Attestations,
				Deposits:          b.Deneb.Message.Body.Deposits,
				VoluntaryExits:    b.Deneb.Message.Body.VoluntaryExits,
				SyncAggregate:     b.Deneb.Message.Body.SyncAggregate,
				ExecutionPayload: &BaseExecutionPayload{
					ParentHash:    b.Deneb.Message.Body.ExecutionPayload.ParentHash,
					FeeRecipient:  b.Deneb.Message.Body.ExecutionPayload.FeeRecipient,
					StateRoot:     b.Deneb.Message.Body.ExecutionPayload.StateRoot,
					ReceiptsRoot:  b.Deneb.Message.Body.ExecutionPayload.ReceiptsRoot,
					LogsBloom:     b.Deneb.Message.Body.ExecutionPayload.LogsBloom,
					PrevRandao:    b.Deneb.Message.Body.ExecutionPayload.PrevRandao,
					BlockNumber:   b.Deneb.Message.Body.ExecutionPayload.BlockNumber,
					GasLimit:      b.Deneb.Message.Body.ExecutionPayload.GasLimit,
					GasUsed:       b.Deneb.Message.Body.ExecutionPayload.GasUsed,
					Timestamp:     b.Deneb.Message.Body.ExecutionPayload.Timestamp,
					ExtraData:     b.Deneb.Message.Body.ExecutionPayload.ExtraData,
					BaseFeePerGas: b.Deneb.Message.Body.ExecutionPayload.BaseFeePerGas,
					BlockHash:     b.Deneb.Message.Body.ExecutionPayload.BlockHash,
					Transactions:  b.Deneb.Message.Body.ExecutionPayload.Transactions,
					Withdrawals:   b.Deneb.Message.Body.ExecutionPayload.Withdrawals,
				},
				BLSToExecutionChanges: b.Deneb.Message.Body.BLSToExecutionChanges,
				BlobKzgCommitments:    b.Deneb.Message.Body.BlobKzgCommitments,
			},
		}
		res.Signature = b.Deneb.Signature
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (b *VersionedSignedBeaconBlock) Version() (string, error) {
	switch {
	case b.Bellatrix != nil:
		return spec.DataVersionBellatrix.String(), nil
	case b.Capella != nil:
		return spec.DataVersionCapella.String(), nil
	case b.Deneb != nil:
		return spec.DataVersionDeneb.String(), nil
	default:
		return "", errors.New("no fork version set")
	}
}

func (b *VersionedSignedBeaconBlock) VersionNumber() (uint64, error) {
	switch {
	case b.Bellatrix != nil:
		return uint64(spec.DataVersionBellatrix), nil
	case b.Capella != nil:
		return uint64(spec.DataVersionCapella), nil
	case b.Deneb != nil:
		return uint64(spec.DataVersionDeneb), nil
	default:
		return 0, errors.New("no fork version set")
	}
}

func (b *VersionedSignedBeaconBlock) WithVersionNumber() (VersionedSignedBeaconBlockWithVersionNumber, error) {
	res := VersionedSignedBeaconBlockWithVersionNumber{}

	versionNumber, err := b.VersionNumber()
	if err != nil {
		return res, err
	}

	res.VersionNumber = versionNumber
	res.VersionedSignedBeaconBlock = b

	return res, nil
}

func (b *VersionedSignedBeaconBlock) WithVersionName() (VersionedSignedBeaconBlockWithVersionName, error) {
	res := VersionedSignedBeaconBlockWithVersionName{}

	versionName, err := b.Version()
	if err != nil {
		return res, err
	}

	res.VersionName = versionName
	res.VersionedSignedBeaconBlock = b

	return res, nil
}
