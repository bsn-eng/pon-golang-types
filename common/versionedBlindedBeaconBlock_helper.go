package common

import (
	"errors"
	"math/big"

	bellatrixApi "github.com/attestantio/go-eth2-client/api/v1/bellatrix"
	capellaApi "github.com/attestantio/go-eth2-client/api/v1/capella"
	denebApi "github.com/attestantio/go-eth2-client/api/v1/deneb"

	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"

	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"

	"github.com/attestantio/go-eth2-client/spec/altair"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseBlindedBeaconBlock struct {
	Slot          phase0.Slot
	ProposerIndex phase0.ValidatorIndex
	ParentRoot    phase0.Root `ssz-size:"32"`
	StateRoot     phase0.Root `ssz-size:"32"`
	Body          *BaseBlindedBeaconBlockBody
}

type BaseBlindedBeaconBlockBody struct {
	RANDAOReveal           phase0.BLSSignature `ssz-size:"96"`
	ETH1Data               *phase0.ETH1Data
	Graffiti               [32]byte                      `ssz-size:"32"`
	ProposerSlashings      []*phase0.ProposerSlashing    `ssz-max:"16"`
	AttesterSlashings      []*phase0.AttesterSlashing    `ssz-max:"2"`
	Attestations           []*phase0.Attestation         `ssz-max:"128"`
	Deposits               []*phase0.Deposit             `ssz-max:"16"`
	VoluntaryExits         []*phase0.SignedVoluntaryExit `ssz-max:"16"`
	SyncAggregate          *altair.SyncAggregate
	ExecutionPayloadHeader *BaseExecutionPayloadHeader
	BLSToExecutionChanges  []*capella.SignedBLSToExecutionChange `ssz-max:"16"`
	BlobKzgCommitments     []deneb.KzgCommitment                 `ssz-max:"4096" ssz-size:"?,48"`
}

func ConstructBlindedBeaconBlock(
	forkVersion string,
	blindedBeaconBlock BaseBlindedBeaconBlock,
) (VersionedBlindedBeaconBlock, error) {

	res := VersionedBlindedBeaconBlock{}

	versionedExecutionPayloadHeader, err := ConstructExecutionPayloadHeader(forkVersion, *blindedBeaconBlock.Body.ExecutionPayloadHeader)
	if err != nil {
		return res, err
	}

	switch forkVersion {
	case "bellatrix":
		res.Bellatrix = &bellatrixApi.BlindedBeaconBlock{
			Slot:          blindedBeaconBlock.Slot,
			ProposerIndex: blindedBeaconBlock.ProposerIndex,
			ParentRoot:    blindedBeaconBlock.ParentRoot,
			StateRoot:     blindedBeaconBlock.StateRoot,
			Body: &bellatrixApi.BlindedBeaconBlockBody{
				RANDAOReveal:           blindedBeaconBlock.Body.RANDAOReveal,
				ETH1Data:               blindedBeaconBlock.Body.ETH1Data,
				Graffiti:               blindedBeaconBlock.Body.Graffiti,
				ProposerSlashings:      blindedBeaconBlock.Body.ProposerSlashings,
				AttesterSlashings:      blindedBeaconBlock.Body.AttesterSlashings,
				Attestations:           blindedBeaconBlock.Body.Attestations,
				Deposits:               blindedBeaconBlock.Body.Deposits,
				VoluntaryExits:         blindedBeaconBlock.Body.VoluntaryExits,
				SyncAggregate:          blindedBeaconBlock.Body.SyncAggregate,
				ExecutionPayloadHeader: versionedExecutionPayloadHeader.Bellatrix,
			},
		}
	case "capella":
		res.Capella = &capellaApi.BlindedBeaconBlock{
			Slot:          blindedBeaconBlock.Slot,
			ProposerIndex: blindedBeaconBlock.ProposerIndex,
			ParentRoot:    blindedBeaconBlock.ParentRoot,
			StateRoot:     blindedBeaconBlock.StateRoot,
			Body: &capellaApi.BlindedBeaconBlockBody{
				RANDAOReveal:           blindedBeaconBlock.Body.RANDAOReveal,
				ETH1Data:               blindedBeaconBlock.Body.ETH1Data,
				Graffiti:               blindedBeaconBlock.Body.Graffiti,
				ProposerSlashings:      blindedBeaconBlock.Body.ProposerSlashings,
				AttesterSlashings:      blindedBeaconBlock.Body.AttesterSlashings,
				Attestations:           blindedBeaconBlock.Body.Attestations,
				Deposits:               blindedBeaconBlock.Body.Deposits,
				VoluntaryExits:         blindedBeaconBlock.Body.VoluntaryExits,
				SyncAggregate:          blindedBeaconBlock.Body.SyncAggregate,
				ExecutionPayloadHeader: versionedExecutionPayloadHeader.Capella,
				BLSToExecutionChanges:  blindedBeaconBlock.Body.BLSToExecutionChanges,
			},
		}
	case "deneb":
		res.Deneb = &denebApi.BlindedBeaconBlock{
			Slot:          blindedBeaconBlock.Slot,
			ProposerIndex: blindedBeaconBlock.ProposerIndex,
			ParentRoot:    blindedBeaconBlock.ParentRoot,
			StateRoot:     blindedBeaconBlock.StateRoot,
			Body: &denebApi.BlindedBeaconBlockBody{
				RANDAOReveal:           blindedBeaconBlock.Body.RANDAOReveal,
				ETH1Data:               blindedBeaconBlock.Body.ETH1Data,
				Graffiti:               blindedBeaconBlock.Body.Graffiti,
				ProposerSlashings:      blindedBeaconBlock.Body.ProposerSlashings,
				AttesterSlashings:      blindedBeaconBlock.Body.AttesterSlashings,
				Attestations:           blindedBeaconBlock.Body.Attestations,
				Deposits:               blindedBeaconBlock.Body.Deposits,
				VoluntaryExits:         blindedBeaconBlock.Body.VoluntaryExits,
				SyncAggregate:          blindedBeaconBlock.Body.SyncAggregate,
				ExecutionPayloadHeader: versionedExecutionPayloadHeader.Deneb,
				BLSToExecutionChanges:  blindedBeaconBlock.Body.BLSToExecutionChanges,
				BlobKzgCommitments:     blindedBeaconBlock.Body.BlobKzgCommitments,
			},
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (b *VersionedBlindedBeaconBlock) ToBaseBlindedBeaconBlock() (BaseBlindedBeaconBlock, error) {
	res := BaseBlindedBeaconBlock{}

	switch {
	case b.Bellatrix != nil:

		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBigInt.SetBytes(b.Capella.Body.ExecutionPayloadHeader.BaseFeePerGas[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.Slot = b.Bellatrix.Slot
		res.ProposerIndex = b.Bellatrix.ProposerIndex
		res.ParentRoot = b.Bellatrix.ParentRoot
		res.StateRoot = b.Bellatrix.StateRoot
		res.Body = &BaseBlindedBeaconBlockBody{
			RANDAOReveal:           b.Bellatrix.Body.RANDAOReveal,
			ETH1Data:               b.Bellatrix.Body.ETH1Data,
			Graffiti:               b.Bellatrix.Body.Graffiti,
			ProposerSlashings:      b.Bellatrix.Body.ProposerSlashings,
			AttesterSlashings:      b.Bellatrix.Body.AttesterSlashings,
			Attestations:           b.Bellatrix.Body.Attestations,
			Deposits:               b.Bellatrix.Body.Deposits,
			VoluntaryExits:         b.Bellatrix.Body.VoluntaryExits,
			SyncAggregate:          b.Bellatrix.Body.SyncAggregate,
			ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
				ParentHash: b.Bellatrix.Body.ExecutionPayloadHeader.ParentHash,
				FeeRecipient: b.Bellatrix.Body.ExecutionPayloadHeader.FeeRecipient,
				StateRoot: b.Bellatrix.Body.ExecutionPayloadHeader.StateRoot,
				ReceiptsRoot: b.Bellatrix.Body.ExecutionPayloadHeader.ReceiptsRoot,
				LogsBloom: b.Bellatrix.Body.ExecutionPayloadHeader.LogsBloom,
				PrevRandao: b.Bellatrix.Body.ExecutionPayloadHeader.PrevRandao,
				BlockNumber: b.Bellatrix.Body.ExecutionPayloadHeader.BlockNumber,
				GasLimit: b.Bellatrix.Body.ExecutionPayloadHeader.GasLimit,
				GasUsed: b.Bellatrix.Body.ExecutionPayloadHeader.GasUsed,
				Timestamp: b.Bellatrix.Body.ExecutionPayloadHeader.Timestamp,
				ExtraData: b.Bellatrix.Body.ExecutionPayloadHeader.ExtraData,
				BaseFeePerGas: baseFeePerGas,
				BlockHash: b.Bellatrix.Body.ExecutionPayloadHeader.BlockHash,
				TransactionsRoot: b.Bellatrix.Body.ExecutionPayloadHeader.TransactionsRoot,
			},
		}
	case b.Capella != nil:

		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBigInt.SetBytes(b.Capella.Body.ExecutionPayloadHeader.BaseFeePerGas[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.Slot = b.Capella.Slot
		res.ProposerIndex = b.Capella.ProposerIndex
		res.ParentRoot = b.Capella.ParentRoot
		res.StateRoot = b.Capella.StateRoot
		res.Body = &BaseBlindedBeaconBlockBody{
			RANDAOReveal:           b.Capella.Body.RANDAOReveal,
			ETH1Data:               b.Capella.Body.ETH1Data,
			Graffiti:               b.Capella.Body.Graffiti,
			ProposerSlashings:      b.Capella.Body.ProposerSlashings,
			AttesterSlashings:      b.Capella.Body.AttesterSlashings,
			Attestations:           b.Capella.Body.Attestations,
			Deposits:               b.Capella.Body.Deposits,
			VoluntaryExits:         b.Capella.Body.VoluntaryExits,
			SyncAggregate:          b.Capella.Body.SyncAggregate,
			ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
				ParentHash: b.Capella.Body.ExecutionPayloadHeader.ParentHash,
				FeeRecipient: b.Capella.Body.ExecutionPayloadHeader.FeeRecipient,
				StateRoot: b.Capella.Body.ExecutionPayloadHeader.StateRoot,
				ReceiptsRoot: b.Capella.Body.ExecutionPayloadHeader.ReceiptsRoot,
				LogsBloom: b.Capella.Body.ExecutionPayloadHeader.LogsBloom,
				PrevRandao: b.Capella.Body.ExecutionPayloadHeader.PrevRandao,
				BlockNumber: b.Capella.Body.ExecutionPayloadHeader.BlockNumber,
				GasLimit: b.Capella.Body.ExecutionPayloadHeader.GasLimit,
				GasUsed: b.Capella.Body.ExecutionPayloadHeader.GasUsed,
				Timestamp: b.Capella.Body.ExecutionPayloadHeader.Timestamp,
				ExtraData: b.Capella.Body.ExecutionPayloadHeader.ExtraData,
				BaseFeePerGas: baseFeePerGas,
				BlockHash: b.Capella.Body.ExecutionPayloadHeader.BlockHash,
				TransactionsRoot: b.Capella.Body.ExecutionPayloadHeader.TransactionsRoot,
				WithdrawalsRoot: b.Capella.Body.ExecutionPayloadHeader.WithdrawalsRoot,
			},
			BLSToExecutionChanges: b.Capella.Body.BLSToExecutionChanges,
		}
	case b.Deneb != nil:
		res.Slot = b.Deneb.Slot
		res.ProposerIndex = b.Deneb.ProposerIndex
		res.ParentRoot = b.Deneb.ParentRoot
		res.StateRoot = b.Deneb.StateRoot
		res.Body = &BaseBlindedBeaconBlockBody{
			RANDAOReveal:           b.Deneb.Body.RANDAOReveal,
			ETH1Data:               b.Deneb.Body.ETH1Data,
			Graffiti:               b.Deneb.Body.Graffiti,
			ProposerSlashings:      b.Deneb.Body.ProposerSlashings,
			AttesterSlashings:      b.Deneb.Body.AttesterSlashings,
			Attestations:           b.Deneb.Body.Attestations,
			Deposits:               b.Deneb.Body.Deposits,
			VoluntaryExits:         b.Deneb.Body.VoluntaryExits,
			SyncAggregate:          b.Deneb.Body.SyncAggregate,
			ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
				ParentHash: b.Deneb.Body.ExecutionPayloadHeader.ParentHash,
				FeeRecipient: b.Deneb.Body.ExecutionPayloadHeader.FeeRecipient,
				StateRoot: b.Deneb.Body.ExecutionPayloadHeader.StateRoot,
				ReceiptsRoot: b.Deneb.Body.ExecutionPayloadHeader.ReceiptsRoot,
				LogsBloom: b.Deneb.Body.ExecutionPayloadHeader.LogsBloom,
				PrevRandao: b.Deneb.Body.ExecutionPayloadHeader.PrevRandao,
				BlockNumber: b.Deneb.Body.ExecutionPayloadHeader.BlockNumber,
				GasLimit: b.Deneb.Body.ExecutionPayloadHeader.GasLimit,
				GasUsed: b.Deneb.Body.ExecutionPayloadHeader.GasUsed,
				Timestamp: b.Deneb.Body.ExecutionPayloadHeader.Timestamp,
				ExtraData: b.Deneb.Body.ExecutionPayloadHeader.ExtraData,
				BaseFeePerGas: b.Deneb.Body.ExecutionPayloadHeader.BaseFeePerGas,
				BlockHash: b.Deneb.Body.ExecutionPayloadHeader.BlockHash,
				TransactionsRoot: b.Deneb.Body.ExecutionPayloadHeader.TransactionsRoot,
				WithdrawalsRoot: b.Deneb.Body.ExecutionPayloadHeader.WithdrawalsRoot,
			},
			BLSToExecutionChanges: b.Deneb.Body.BLSToExecutionChanges,
			BlobKzgCommitments: b.Deneb.Body.BlobKzgCommitments,
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil

}

// Converts the VersionedBeaconBlock to a VersionedBeaconBlock without the transactions and withdrawals as those cannot be derived
func (b *VersionedBlindedBeaconBlock) ToVersionedBeaconBlock() (VersionedBeaconBlock, error) {
	res := VersionedBeaconBlock{}

	baseBlindedBeaconBlock, err := b.ToBaseBlindedBeaconBlock()
	if err != nil {
		return res, err
	}

	baseBeaconBlock := BaseBeaconBlock{
		Slot: 		baseBlindedBeaconBlock.Slot,
		ProposerIndex: 	baseBlindedBeaconBlock.ProposerIndex,
		ParentRoot: 	baseBlindedBeaconBlock.ParentRoot,
		StateRoot: 	baseBlindedBeaconBlock.StateRoot,
		Body: &BaseBeaconBlockBody{
			RANDAOReveal:           baseBlindedBeaconBlock.Body.RANDAOReveal,
			ETH1Data:               baseBlindedBeaconBlock.Body.ETH1Data,
			Graffiti:               baseBlindedBeaconBlock.Body.Graffiti,
			ProposerSlashings:      baseBlindedBeaconBlock.Body.ProposerSlashings,
			AttesterSlashings:      baseBlindedBeaconBlock.Body.AttesterSlashings,
			Attestations:           baseBlindedBeaconBlock.Body.Attestations,
			Deposits:               baseBlindedBeaconBlock.Body.Deposits,
			VoluntaryExits:         baseBlindedBeaconBlock.Body.VoluntaryExits,
			SyncAggregate:          baseBlindedBeaconBlock.Body.SyncAggregate,
			ExecutionPayload: &BaseExecutionPayload{
				ParentHash: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.ParentHash,
				FeeRecipient: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.FeeRecipient,
				StateRoot: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.StateRoot,
				ReceiptsRoot: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.ReceiptsRoot,
				LogsBloom: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.LogsBloom,
				PrevRandao: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.PrevRandao,
				BlockNumber: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.BlockNumber,
				GasLimit: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.GasLimit,
				GasUsed: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.GasUsed,
				Timestamp: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.Timestamp,
				ExtraData: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.ExtraData,
				BaseFeePerGas: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.BaseFeePerGas,
				BlockHash: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.BlockHash,
				// Transactions:
				// Withdrawals:
				BlobGasUsed: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.BlobGasUsed,
				ExcessBlobGas: baseBlindedBeaconBlock.Body.ExecutionPayloadHeader.ExcessBlobGas,
			},
			BLSToExecutionChanges: baseBlindedBeaconBlock.Body.BLSToExecutionChanges,
			BlobKzgCommitments: baseBlindedBeaconBlock.Body.BlobKzgCommitments,
		},
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

	res, err = ConstructBeaconBlock(forkVersion, baseBeaconBlock)
	if err != nil {
		return res, err
	}

	return res, nil
}