package common

import (
	"errors"
	"math/big"

	"github.com/attestantio/go-eth2-client/spec"

	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"

	"github.com/attestantio/go-eth2-client/spec/altair"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseBeaconBlock struct {
	Slot          phase0.Slot
	ProposerIndex phase0.ValidatorIndex
	ParentRoot    phase0.Root `ssz-size:"32"`
	StateRoot     phase0.Root `ssz-size:"32"`
	Body          *BaseBeaconBlockBody
}

type BaseBeaconBlockBody struct {
	RANDAOReveal          phase0.BLSSignature `ssz-size:"96"`
	ETH1Data              *phase0.ETH1Data
	Graffiti              [32]byte                      `ssz-size:"32"`
	ProposerSlashings     []*phase0.ProposerSlashing    `ssz-max:"16"`
	AttesterSlashings     []*phase0.AttesterSlashing    `ssz-max:"2"`
	Attestations          []*phase0.Attestation         `ssz-max:"128"`
	Deposits              []*phase0.Deposit             `ssz-max:"16"`
	VoluntaryExits        []*phase0.SignedVoluntaryExit `ssz-max:"16"`
	SyncAggregate         *altair.SyncAggregate
	ExecutionPayload      *BaseExecutionPayload
	BLSToExecutionChanges []*capella.SignedBLSToExecutionChange `ssz-max:"16"`
	BlobKzgCommitments    []deneb.KzgCommitment                 `ssz-max:"4096" ssz-size:"?,48"`
}

func ConstructBeaconBlock(
	forkVersion string,
	beaconBlock BaseBeaconBlock,
) (VersionedBeaconBlock, error) {

	res := VersionedBeaconBlock{}

	versionedExecutionPayload, err := ConstructExecutionPayload(forkVersion, *beaconBlock.Body.ExecutionPayload)
	if err != nil {
		return res, err
	}

	switch forkVersion {
	case spec.DataVersionBellatrix.String():
		res.Bellatrix = &bellatrix.BeaconBlock{
			Slot:          beaconBlock.Slot,
			ProposerIndex: beaconBlock.ProposerIndex,
			ParentRoot:    beaconBlock.ParentRoot,
			StateRoot:     beaconBlock.StateRoot,
			Body: &bellatrix.BeaconBlockBody{
				RANDAOReveal:      beaconBlock.Body.RANDAOReveal,
				ETH1Data:          beaconBlock.Body.ETH1Data,
				Graffiti:          beaconBlock.Body.Graffiti,
				ProposerSlashings: beaconBlock.Body.ProposerSlashings,
				AttesterSlashings: beaconBlock.Body.AttesterSlashings,
				Attestations:      beaconBlock.Body.Attestations,
				Deposits:          beaconBlock.Body.Deposits,
				VoluntaryExits:    beaconBlock.Body.VoluntaryExits,
				SyncAggregate:     beaconBlock.Body.SyncAggregate,
				ExecutionPayload:  versionedExecutionPayload.Bellatrix,
			},
		}
	case spec.DataVersionCapella.String():
		res.Capella = &capella.BeaconBlock{
			Slot:          beaconBlock.Slot,
			ProposerIndex: beaconBlock.ProposerIndex,
			ParentRoot:    beaconBlock.ParentRoot,
			StateRoot:     beaconBlock.StateRoot,
			Body: &capella.BeaconBlockBody{
				RANDAOReveal:          beaconBlock.Body.RANDAOReveal,
				ETH1Data:              beaconBlock.Body.ETH1Data,
				Graffiti:              beaconBlock.Body.Graffiti,
				ProposerSlashings:     beaconBlock.Body.ProposerSlashings,
				AttesterSlashings:     beaconBlock.Body.AttesterSlashings,
				Attestations:          beaconBlock.Body.Attestations,
				Deposits:              beaconBlock.Body.Deposits,
				VoluntaryExits:        beaconBlock.Body.VoluntaryExits,
				SyncAggregate:         beaconBlock.Body.SyncAggregate,
				ExecutionPayload:      versionedExecutionPayload.Capella,
				BLSToExecutionChanges: beaconBlock.Body.BLSToExecutionChanges,
			},
		}
	case spec.DataVersionDeneb.String():
		res.Deneb = &deneb.BeaconBlock{
			Slot:          beaconBlock.Slot,
			ProposerIndex: beaconBlock.ProposerIndex,
			ParentRoot:    beaconBlock.ParentRoot,
			StateRoot:     beaconBlock.StateRoot,
			Body: &deneb.BeaconBlockBody{
				RANDAOReveal:          beaconBlock.Body.RANDAOReveal,
				ETH1Data:              beaconBlock.Body.ETH1Data,
				Graffiti:              beaconBlock.Body.Graffiti,
				ProposerSlashings:     beaconBlock.Body.ProposerSlashings,
				AttesterSlashings:     beaconBlock.Body.AttesterSlashings,
				Attestations:          beaconBlock.Body.Attestations,
				Deposits:              beaconBlock.Body.Deposits,
				VoluntaryExits:        beaconBlock.Body.VoluntaryExits,
				SyncAggregate:         beaconBlock.Body.SyncAggregate,
				ExecutionPayload:      versionedExecutionPayload.Deneb,
				BLSToExecutionChanges: beaconBlock.Body.BLSToExecutionChanges,
				BlobKzgCommitments:    beaconBlock.Body.BlobKzgCommitments,
			},
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (b *VersionedBeaconBlock) ToBaseBeaconBlock() (BaseBeaconBlock, error) {
	res := BaseBeaconBlock{}

	switch {
	case b.Bellatrix != nil:

		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Bellatrix.Body.ExecutionPayload.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Bellatrix.Body.ExecutionPayload.BaseFeePerGas[len(b.Bellatrix.Body.ExecutionPayload.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.ParentRoot = b.Bellatrix.ParentRoot
		res.ProposerIndex = b.Bellatrix.ProposerIndex
		res.Slot = b.Bellatrix.Slot
		res.StateRoot = b.Bellatrix.StateRoot
		res.Body = &BaseBeaconBlockBody{
			RANDAOReveal:      b.Bellatrix.Body.RANDAOReveal,
			ETH1Data:          b.Bellatrix.Body.ETH1Data,
			Graffiti:          b.Bellatrix.Body.Graffiti,
			ProposerSlashings: b.Bellatrix.Body.ProposerSlashings,
			AttesterSlashings: b.Bellatrix.Body.AttesterSlashings,
			Attestations:      b.Bellatrix.Body.Attestations,
			Deposits:          b.Bellatrix.Body.Deposits,
			VoluntaryExits:    b.Bellatrix.Body.VoluntaryExits,
			SyncAggregate:     b.Bellatrix.Body.SyncAggregate,
			ExecutionPayload: &BaseExecutionPayload{
				ParentHash:    b.Bellatrix.Body.ExecutionPayload.ParentHash,
				FeeRecipient:  b.Bellatrix.Body.ExecutionPayload.FeeRecipient,
				StateRoot:     b.Bellatrix.Body.ExecutionPayload.StateRoot,
				ReceiptsRoot:  b.Bellatrix.Body.ExecutionPayload.ReceiptsRoot,
				LogsBloom:     b.Bellatrix.Body.ExecutionPayload.LogsBloom,
				PrevRandao:    b.Bellatrix.Body.ExecutionPayload.PrevRandao,
				BlockNumber:   b.Bellatrix.Body.ExecutionPayload.BlockNumber,
				GasLimit:      b.Bellatrix.Body.ExecutionPayload.GasLimit,
				GasUsed:       b.Bellatrix.Body.ExecutionPayload.GasUsed,
				Timestamp:     b.Bellatrix.Body.ExecutionPayload.Timestamp,
				ExtraData:     b.Bellatrix.Body.ExecutionPayload.ExtraData,
				BaseFeePerGas: baseFeePerGas,
				BlockHash:     b.Bellatrix.Body.ExecutionPayload.BlockHash,
				Transactions:  b.Bellatrix.Body.ExecutionPayload.Transactions,
			},
		}

	case b.Capella != nil:

		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(b.Capella.Body.ExecutionPayload.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = b.Capella.Body.ExecutionPayload.BaseFeePerGas[len(b.Capella.Body.ExecutionPayload.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}

		res.ParentRoot = b.Capella.ParentRoot
		res.ProposerIndex = b.Capella.ProposerIndex
		res.Slot = b.Capella.Slot
		res.StateRoot = b.Capella.StateRoot
		res.Body = &BaseBeaconBlockBody{
			RANDAOReveal:      b.Capella.Body.RANDAOReveal,
			ETH1Data:          b.Capella.Body.ETH1Data,
			Graffiti:          b.Capella.Body.Graffiti,
			ProposerSlashings: b.Capella.Body.ProposerSlashings,
			AttesterSlashings: b.Capella.Body.AttesterSlashings,
			Attestations:      b.Capella.Body.Attestations,
			Deposits:          b.Capella.Body.Deposits,
			VoluntaryExits:    b.Capella.Body.VoluntaryExits,
			SyncAggregate:     b.Capella.Body.SyncAggregate,
			ExecutionPayload: &BaseExecutionPayload{
				ParentHash:    b.Capella.Body.ExecutionPayload.ParentHash,
				FeeRecipient:  b.Capella.Body.ExecutionPayload.FeeRecipient,
				StateRoot:     b.Capella.Body.ExecutionPayload.StateRoot,
				ReceiptsRoot:  b.Capella.Body.ExecutionPayload.ReceiptsRoot,
				LogsBloom:     b.Capella.Body.ExecutionPayload.LogsBloom,
				PrevRandao:    b.Capella.Body.ExecutionPayload.PrevRandao,
				BlockNumber:   b.Capella.Body.ExecutionPayload.BlockNumber,
				GasLimit:      b.Capella.Body.ExecutionPayload.GasLimit,
				GasUsed:       b.Capella.Body.ExecutionPayload.GasUsed,
				Timestamp:     b.Capella.Body.ExecutionPayload.Timestamp,
				ExtraData:     b.Capella.Body.ExecutionPayload.ExtraData,
				BaseFeePerGas: baseFeePerGas,
				BlockHash:     b.Capella.Body.ExecutionPayload.BlockHash,
				Transactions:  b.Capella.Body.ExecutionPayload.Transactions,
				Withdrawals:   b.Capella.Body.ExecutionPayload.Withdrawals,
			},
			BLSToExecutionChanges: b.Capella.Body.BLSToExecutionChanges,
		}

	case b.Deneb != nil:

		res.ParentRoot = b.Deneb.ParentRoot
		res.ProposerIndex = b.Deneb.ProposerIndex
		res.Slot = b.Deneb.Slot
		res.StateRoot = b.Deneb.StateRoot
		res.Body = &BaseBeaconBlockBody{
			RANDAOReveal:      b.Deneb.Body.RANDAOReveal,
			ETH1Data:          b.Deneb.Body.ETH1Data,
			Graffiti:          b.Deneb.Body.Graffiti,
			ProposerSlashings: b.Deneb.Body.ProposerSlashings,
			AttesterSlashings: b.Deneb.Body.AttesterSlashings,
			Attestations:      b.Deneb.Body.Attestations,
			Deposits:          b.Deneb.Body.Deposits,
			VoluntaryExits:    b.Deneb.Body.VoluntaryExits,
			SyncAggregate:     b.Deneb.Body.SyncAggregate,
			ExecutionPayload: &BaseExecutionPayload{
				ParentHash:    b.Deneb.Body.ExecutionPayload.ParentHash,
				FeeRecipient:  b.Deneb.Body.ExecutionPayload.FeeRecipient,
				StateRoot:     b.Deneb.Body.ExecutionPayload.StateRoot,
				ReceiptsRoot:  b.Deneb.Body.ExecutionPayload.ReceiptsRoot,
				LogsBloom:     b.Deneb.Body.ExecutionPayload.LogsBloom,
				PrevRandao:    b.Deneb.Body.ExecutionPayload.PrevRandao,
				BlockNumber:   b.Deneb.Body.ExecutionPayload.BlockNumber,
				GasLimit:      b.Deneb.Body.ExecutionPayload.GasLimit,
				GasUsed:       b.Deneb.Body.ExecutionPayload.GasUsed,
				Timestamp:     b.Deneb.Body.ExecutionPayload.Timestamp,
				ExtraData:     b.Deneb.Body.ExecutionPayload.ExtraData,
				BaseFeePerGas: b.Deneb.Body.ExecutionPayload.BaseFeePerGas,
				BlockHash:     b.Deneb.Body.ExecutionPayload.BlockHash,
				Transactions:  b.Deneb.Body.ExecutionPayload.Transactions,
				Withdrawals:   b.Deneb.Body.ExecutionPayload.Withdrawals,
			},
			BLSToExecutionChanges: b.Deneb.Body.BLSToExecutionChanges,
			BlobKzgCommitments:    b.Deneb.Body.BlobKzgCommitments,
		}
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

// Converts the VersionedBeaconBlock to a VersionedBlindedBeaconBlock, hasihing the transactions and withdrawals
func (b *VersionedBeaconBlock) ToVersionedBlindedBeaconBlock() (VersionedBlindedBeaconBlock, error) {
	res := VersionedBlindedBeaconBlock{}

	baseBeaconBlock, err := b.ToBaseBeaconBlock()
	if err != nil {
		return res, err
	}

	transactionsRoot, err := ComputeTransactionsRoot(baseBeaconBlock.Body.ExecutionPayload.Transactions)
	if err != nil {
		return res, err
	}

	withdrawalsRoot, err := ComputeWithdrawalsRoot(baseBeaconBlock.Body.ExecutionPayload.Withdrawals)
	if err != nil {
		return res, err
	}

	baseBlindedBeaconBlock := BaseBlindedBeaconBlock{
		Slot:          baseBeaconBlock.Slot,
		ProposerIndex: baseBeaconBlock.ProposerIndex,
		ParentRoot:    baseBeaconBlock.ParentRoot,
		StateRoot:     baseBeaconBlock.StateRoot,
		Body: &BaseBlindedBeaconBlockBody{
			RANDAOReveal:      baseBeaconBlock.Body.RANDAOReveal,
			ETH1Data:          baseBeaconBlock.Body.ETH1Data,
			Graffiti:          baseBeaconBlock.Body.Graffiti,
			ProposerSlashings: baseBeaconBlock.Body.ProposerSlashings,
			AttesterSlashings: baseBeaconBlock.Body.AttesterSlashings,
			Attestations:      baseBeaconBlock.Body.Attestations,
			Deposits:          baseBeaconBlock.Body.Deposits,
			VoluntaryExits:    baseBeaconBlock.Body.VoluntaryExits,
			SyncAggregate:     baseBeaconBlock.Body.SyncAggregate,
			ExecutionPayloadHeader: &BaseExecutionPayloadHeader{
				ParentHash:       baseBeaconBlock.Body.ExecutionPayload.ParentHash,
				FeeRecipient:     baseBeaconBlock.Body.ExecutionPayload.FeeRecipient,
				StateRoot:        baseBeaconBlock.Body.ExecutionPayload.StateRoot,
				ReceiptsRoot:     baseBeaconBlock.Body.ExecutionPayload.ReceiptsRoot,
				LogsBloom:        baseBeaconBlock.Body.ExecutionPayload.LogsBloom,
				PrevRandao:       baseBeaconBlock.Body.ExecutionPayload.PrevRandao,
				BlockNumber:      baseBeaconBlock.Body.ExecutionPayload.BlockNumber,
				GasLimit:         baseBeaconBlock.Body.ExecutionPayload.GasLimit,
				GasUsed:          baseBeaconBlock.Body.ExecutionPayload.GasUsed,
				Timestamp:        baseBeaconBlock.Body.ExecutionPayload.Timestamp,
				ExtraData:        baseBeaconBlock.Body.ExecutionPayload.ExtraData,
				BaseFeePerGas:    baseBeaconBlock.Body.ExecutionPayload.BaseFeePerGas,
				BlockHash:        baseBeaconBlock.Body.ExecutionPayload.BlockHash,
				TransactionsRoot: transactionsRoot,
				WithdrawalsRoot:  withdrawalsRoot,
				BlobGasUsed:      baseBeaconBlock.Body.ExecutionPayload.BlobGasUsed,
				ExcessBlobGas:    baseBeaconBlock.Body.ExecutionPayload.ExcessBlobGas,
			},
			BLSToExecutionChanges: baseBeaconBlock.Body.BLSToExecutionChanges,
			BlobKzgCommitments:    baseBeaconBlock.Body.BlobKzgCommitments,
		},
	}

	var forkVersion string
	switch {
	case b.Bellatrix != nil:
		forkVersion = spec.DataVersionBellatrix.String()
	case b.Capella != nil:
		forkVersion = spec.DataVersionCapella.String()
	case b.Deneb != nil:
		forkVersion = spec.DataVersionDeneb.String()
	default:
		return res, errors.New("unsupported fork version")
	}

	res, err = ConstructBlindedBeaconBlock(forkVersion, baseBlindedBeaconBlock)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *VersionedBeaconBlock) Version() (string, error) {
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

func (b *VersionedBeaconBlock) VersionNumber() (uint64, error) {
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

func (b *VersionedBeaconBlock) WithVersionNumber() (VersionedBeaconBlockWithVersionNumber, error) {
	res := VersionedBeaconBlockWithVersionNumber{}

	versionNumber, err := b.VersionNumber()
	if err != nil {
		return res, err
	}

	res.VersionNumber = versionNumber
	res.VersionedBeaconBlock = b

	return res, nil
}

func (b *VersionedBeaconBlock) WithVersionName() (VersionedBeaconBlockWithVersionName, error) {
	res := VersionedBeaconBlockWithVersionName{}

	version, err := b.Version()
	if err != nil {
		return res, err
	}

	res.VersionName = version
	res.VersionedBeaconBlock = b

	return res, nil
}