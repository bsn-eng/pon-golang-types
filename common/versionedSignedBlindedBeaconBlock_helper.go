package common

import (
	"errors"
	"math/big"

	bellatrix "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	capella "github.com/attestantio/go-eth2-client/api/v1/capella"
	deneb "github.com/attestantio/go-eth2-client/api/v1/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseSignedBlindedBeaconBlock struct {
	Message   *BaseBlindedBeaconBlock
	Signature phase0.BLSSignature `ssz-size:"96"`
}

func ConstructSignedBlindedBeaconBlock(
	forkVersion string,
	signedBlindedBeaconBlock BaseSignedBlindedBeaconBlock,
) (VersionedSignedBlindedBeaconBlock, error) {

	res := VersionedSignedBlindedBeaconBlock{}

	VersionedBlindedBeaconBlock, err := ConstructBlindedBeaconBlock(forkVersion, *signedBlindedBeaconBlock.Message)
	if err != nil {
		return res, err
	}

	switch forkVersion {
	case "bellatrix":
		res.Bellatrix = &bellatrix.SignedBlindedBeaconBlock{
			Message:   VersionedBlindedBeaconBlock.Bellatrix,
			Signature: signedBlindedBeaconBlock.Signature,
		}
	case "capella":
		res.Capella = &capella.SignedBlindedBeaconBlock{
			Message:   VersionedBlindedBeaconBlock.Capella,
			Signature: signedBlindedBeaconBlock.Signature,
		}
	case "deneb":
		res.Deneb = &deneb.SignedBlindedBeaconBlock{
			Message:   VersionedBlindedBeaconBlock.Deneb,
			Signature: signedBlindedBeaconBlock.Signature,
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (b *VersionedSignedBlindedBeaconBlock) ToBaseSignedBlindedBeaconBlock() (BaseSignedBlindedBeaconBlock, error) {
	res := BaseSignedBlindedBeaconBlock{}

	switch {
	case b.Bellatrix != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Bellatrix.Message.Body.ExecutionPayloadHeader.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Bellatrix.Message.Body.ExecutionPayloadHeader.BaseFeePerGas[len(b.Bellatrix.Message.Body.ExecutionPayloadHeader.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}
		res.Message = &BaseBlindedBeaconBlock{
			Slot:          b.Bellatrix.Message.Slot,
			ProposerIndex: b.Bellatrix.Message.ProposerIndex,
			ParentRoot:    b.Bellatrix.Message.ParentRoot,
			StateRoot:     b.Bellatrix.Message.StateRoot,
			Body:          &BaseBlindedBeaconBlockBody{
				RANDAOReveal:      b.Bellatrix.Message.Body.RANDAOReveal,
				ETH1Data:          b.Bellatrix.Message.Body.ETH1Data,
				Graffiti:          b.Bellatrix.Message.Body.Graffiti,
				ProposerSlashings: b.Bellatrix.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Bellatrix.Message.Body.AttesterSlashings,
				Attestations:      b.Bellatrix.Message.Body.Attestations,
				Deposits:          b.Bellatrix.Message.Body.Deposits,
				VoluntaryExits:    b.Bellatrix.Message.Body.VoluntaryExits,
				SyncAggregate:     b.Bellatrix.Message.Body.SyncAggregate,
				ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
					ParentHash:    b.Bellatrix.Message.Body.ExecutionPayloadHeader.ParentHash,
					FeeRecipient:  b.Bellatrix.Message.Body.ExecutionPayloadHeader.FeeRecipient,
					StateRoot:     b.Bellatrix.Message.Body.ExecutionPayloadHeader.StateRoot,
					ReceiptsRoot:  b.Bellatrix.Message.Body.ExecutionPayloadHeader.ReceiptsRoot,
					LogsBloom:     b.Bellatrix.Message.Body.ExecutionPayloadHeader.LogsBloom,
					PrevRandao:    b.Bellatrix.Message.Body.ExecutionPayloadHeader.PrevRandao,
					BlockNumber:   b.Bellatrix.Message.Body.ExecutionPayloadHeader.BlockNumber,
					GasLimit:      b.Bellatrix.Message.Body.ExecutionPayloadHeader.GasLimit,
					GasUsed:       b.Bellatrix.Message.Body.ExecutionPayloadHeader.GasUsed,
					Timestamp:     b.Bellatrix.Message.Body.ExecutionPayloadHeader.Timestamp,
					ExtraData:     b.Bellatrix.Message.Body.ExecutionPayloadHeader.ExtraData,
					BaseFeePerGas: baseFeePerGas,
					BlockHash:     b.Bellatrix.Message.Body.ExecutionPayloadHeader.BlockHash,
					TransactionsRoot: b.Bellatrix.Message.Body.ExecutionPayloadHeader.TransactionsRoot,
				},
			},
		}
		res.Signature = b.Bellatrix.Signature
	case b.Capella != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Capella.Message.Body.ExecutionPayloadHeader.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Capella.Message.Body.ExecutionPayloadHeader.BaseFeePerGas[len(b.Capella.Message.Body.ExecutionPayloadHeader.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.Message = &BaseBlindedBeaconBlock{
			Slot:          b.Capella.Message.Slot,
			ProposerIndex: b.Capella.Message.ProposerIndex,
			ParentRoot:    b.Capella.Message.ParentRoot,
			StateRoot:     b.Capella.Message.StateRoot,
			Body:          &BaseBlindedBeaconBlockBody{
				RANDAOReveal:      b.Capella.Message.Body.RANDAOReveal,
				ETH1Data:          b.Capella.Message.Body.ETH1Data,
				Graffiti:          b.Capella.Message.Body.Graffiti,
				ProposerSlashings: b.Capella.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Capella.Message.Body.AttesterSlashings,
				Attestations:      b.Capella.Message.Body.Attestations,
				Deposits:          b.Capella.Message.Body.Deposits,
				VoluntaryExits:    b.Capella.Message.Body.VoluntaryExits,
				SyncAggregate:     b.Capella.Message.Body.SyncAggregate,
				ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
					ParentHash:    b.Capella.Message.Body.ExecutionPayloadHeader.ParentHash,
					FeeRecipient:  b.Capella.Message.Body.ExecutionPayloadHeader.FeeRecipient,
					StateRoot:     b.Capella.Message.Body.ExecutionPayloadHeader.StateRoot,
					ReceiptsRoot:  b.Capella.Message.Body.ExecutionPayloadHeader.ReceiptsRoot,
					LogsBloom:     b.Capella.Message.Body.ExecutionPayloadHeader.LogsBloom,
					PrevRandao:    b.Capella.Message.Body.ExecutionPayloadHeader.PrevRandao,
					BlockNumber:   b.Capella.Message.Body.ExecutionPayloadHeader.BlockNumber,
					GasLimit:      b.Capella.Message.Body.ExecutionPayloadHeader.GasLimit,
					GasUsed:       b.Capella.Message.Body.ExecutionPayloadHeader.GasUsed,
					Timestamp:     b.Capella.Message.Body.ExecutionPayloadHeader.Timestamp,
					ExtraData:     b.Capella.Message.Body.ExecutionPayloadHeader.ExtraData,
					BaseFeePerGas: baseFeePerGas,
					BlockHash:     b.Capella.Message.Body.ExecutionPayloadHeader.BlockHash,
					TransactionsRoot:  b.Capella.Message.Body.ExecutionPayloadHeader.TransactionsRoot,
					WithdrawalsRoot:  b.Capella.Message.Body.ExecutionPayloadHeader.WithdrawalsRoot,
				},
				BLSToExecutionChanges: b.Capella.Message.Body.BLSToExecutionChanges,
			},
		}
		res.Signature = b.Capella.Signature
	case b.Deneb != nil:

		res.Message = &BaseBlindedBeaconBlock{
			Slot: 		b.Deneb.Message.Slot,
			ProposerIndex: b.Deneb.Message.ProposerIndex,
			ParentRoot: b.Deneb.Message.ParentRoot,
			StateRoot: b.Deneb.Message.StateRoot,
			Body: &BaseBlindedBeaconBlockBody{
				RANDAOReveal: b.Deneb.Message.Body.RANDAOReveal,
				ETH1Data: b.Deneb.Message.Body.ETH1Data,
				Graffiti: b.Deneb.Message.Body.Graffiti,
				ProposerSlashings: b.Deneb.Message.Body.ProposerSlashings,
				AttesterSlashings: b.Deneb.Message.Body.AttesterSlashings,
				Attestations: b.Deneb.Message.Body.Attestations,
				Deposits: b.Deneb.Message.Body.Deposits,
				VoluntaryExits: b.Deneb.Message.Body.VoluntaryExits,
				SyncAggregate: b.Deneb.Message.Body.SyncAggregate,
				ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
					ParentHash: b.Deneb.Message.Body.ExecutionPayloadHeader.ParentHash,
					FeeRecipient: b.Deneb.Message.Body.ExecutionPayloadHeader.FeeRecipient,
					StateRoot: b.Deneb.Message.Body.ExecutionPayloadHeader.StateRoot,
					ReceiptsRoot: b.Deneb.Message.Body.ExecutionPayloadHeader.ReceiptsRoot,
					LogsBloom: b.Deneb.Message.Body.ExecutionPayloadHeader.LogsBloom,
					PrevRandao: b.Deneb.Message.Body.ExecutionPayloadHeader.PrevRandao,
					BlockNumber: b.Deneb.Message.Body.ExecutionPayloadHeader.BlockNumber,
					GasLimit: b.Deneb.Message.Body.ExecutionPayloadHeader.GasLimit,
					GasUsed: b.Deneb.Message.Body.ExecutionPayloadHeader.GasUsed,
					Timestamp: b.Deneb.Message.Body.ExecutionPayloadHeader.Timestamp,
					ExtraData: b.Deneb.Message.Body.ExecutionPayloadHeader.ExtraData,
					BaseFeePerGas: b.Deneb.Message.Body.ExecutionPayloadHeader.BaseFeePerGas,
					BlockHash: b.Deneb.Message.Body.ExecutionPayloadHeader.BlockHash,
					TransactionsRoot: b.Deneb.Message.Body.ExecutionPayloadHeader.TransactionsRoot,
					WithdrawalsRoot: b.Deneb.Message.Body.ExecutionPayloadHeader.WithdrawalsRoot,
				},
				BLSToExecutionChanges: b.Deneb.Message.Body.BLSToExecutionChanges,
				BlobKzgCommitments: b.Deneb.Message.Body.BlobKzgCommitments,
			},
		}
		res.Signature = b.Deneb.Signature
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

// Converts the VersionedSignedBlindedBeaconBlock to a VersionedSignedBeaconBlock without the transactions and withdrawals as those cannot be derived
func (b *VersionedSignedBlindedBeaconBlock) ToVersionedSignedBeaconBlock() (VersionedSignedBeaconBlock, error) {
	res := VersionedSignedBeaconBlock{}

	baseSignedBlindedBeaconBlock, err := b.ToBaseSignedBlindedBeaconBlock()
	if err != nil {
		return res, err
	}

	baseSignedBeaconBlock := BaseSignedBeaconBlock{
		Message:   &BaseBeaconBlock{
			Slot: 		baseSignedBlindedBeaconBlock.Message.Slot,
			ProposerIndex: 	baseSignedBlindedBeaconBlock.Message.ProposerIndex,
			ParentRoot: 	baseSignedBlindedBeaconBlock.Message.ParentRoot,
			StateRoot: 	baseSignedBlindedBeaconBlock.Message.StateRoot,
			Body: &BaseBeaconBlockBody{
				RANDAOReveal:           baseSignedBlindedBeaconBlock.Message.Body.RANDAOReveal,
				ETH1Data:               baseSignedBlindedBeaconBlock.Message.Body.ETH1Data,
				Graffiti:               baseSignedBlindedBeaconBlock.Message.Body.Graffiti,
				ProposerSlashings:      baseSignedBlindedBeaconBlock.Message.Body.ProposerSlashings,
				AttesterSlashings:      baseSignedBlindedBeaconBlock.Message.Body.AttesterSlashings,
				Attestations:           baseSignedBlindedBeaconBlock.Message.Body.Attestations,
				Deposits:               baseSignedBlindedBeaconBlock.Message.Body.Deposits,
				VoluntaryExits:         baseSignedBlindedBeaconBlock.Message.Body.VoluntaryExits,
				SyncAggregate:          baseSignedBlindedBeaconBlock.Message.Body.SyncAggregate,
				ExecutionPayload: &BaseExecutionPayload{
					ParentHash: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.ParentHash,
					FeeRecipient: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.FeeRecipient,
					StateRoot: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.StateRoot,
					ReceiptsRoot: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.ReceiptsRoot,
					LogsBloom: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.LogsBloom,
					PrevRandao: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.PrevRandao,
					BlockNumber: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.BlockNumber,
					GasLimit: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.GasLimit,
					GasUsed: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.GasUsed,
					Timestamp: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.Timestamp,
					ExtraData: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.ExtraData,
					BaseFeePerGas: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.BaseFeePerGas,
					BlockHash: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.BlockHash,
					// Transactions:
					// Withdrawals:
					BlobGasUsed: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.BlobGasUsed,
					ExcessBlobGas: baseSignedBlindedBeaconBlock.Message.Body.ExecutionPayloadHeader.ExcessBlobGas,
				},
				BLSToExecutionChanges: baseSignedBlindedBeaconBlock.Message.Body.BLSToExecutionChanges,
				BlobKzgCommitments: baseSignedBlindedBeaconBlock.Message.Body.BlobKzgCommitments,
			},
		},
		Signature: baseSignedBlindedBeaconBlock.Signature,
	}

	var forkVersion string
	switch {
	case b.Bellatrix != nil:
		forkVersion = "bellatrix"
	case b.Capella != nil:
		forkVersion = "capella"
	case b.Deneb != nil:
		forkVersion = "deneb"
	default:
		return res, errors.New("unsupported fork version")
	}

	res, err = ConstructSignedBeaconBlock(forkVersion, baseSignedBeaconBlock)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *VersionedSignedBlindedBeaconBlock) Version() (string, error) {
	switch {
	case b.Bellatrix != nil:
		return "bellatrix", nil
	case b.Capella != nil:
		return "capella", nil
	case b.Deneb != nil:
		return "deneb", nil
	default:
		return "", errors.New("no fork version set")
	}
}