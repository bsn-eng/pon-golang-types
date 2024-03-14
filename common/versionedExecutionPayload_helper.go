package common

import (
	"errors"
	"math/big"

	"github.com/attestantio/go-eth2-client/spec"

	denebApi "github.com/attestantio/go-builder-client/api/deneb"
	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseExecutionPayload struct {
	ParentHash    phase0.Hash32              `ssz-size:"32"`
	FeeRecipient  bellatrix.ExecutionAddress `ssz-size:"20"`
	StateRoot     phase0.Root                `ssz-size:"32"`
	ReceiptsRoot  phase0.Root                `ssz-size:"32"`
	LogsBloom     [256]byte                  `ssz-size:"256"`
	PrevRandao    [32]byte                   `ssz-size:"32"`
	BlockNumber   uint64
	GasLimit      uint64
	GasUsed       uint64
	Timestamp     uint64
	ExtraData     []byte                  `ssz-max:"32"`
	BaseFeePerGas *uint256.Int            `ssz-size:"32"`
	BlockHash     phase0.Hash32           `ssz-size:"32"`
	Transactions  []bellatrix.Transaction `ssz-max:"1048576,1073741824" ssz-size:"?,?"`
	Withdrawals   []*capella.Withdrawal   `ssz-max:"16"`
	BlobGasUsed   uint64
	ExcessBlobGas uint64
	BlobsBundle   *denebApi.BlobsBundle
}

func ConstructExecutionPayload(
	forkVersion string,
	executionPayload BaseExecutionPayload,
) (VersionedExecutionPayload, error) {

	res := VersionedExecutionPayload{}

	switch forkVersion {
	case spec.DataVersionBellatrix.String():

		baseFeePerGas := [32]byte(executionPayload.BaseFeePerGas.PaddedBytes(32))
		baseFeePerGasLE := [32]byte{}
		for i := 0; i < len(baseFeePerGas); i++ {
			baseFeePerGasLE[i] = baseFeePerGas[len(baseFeePerGas)-1-i]
		}

		res.Bellatrix = &bellatrix.ExecutionPayload{
			ParentHash:    executionPayload.ParentHash,
			FeeRecipient:  executionPayload.FeeRecipient,
			StateRoot:     executionPayload.StateRoot,
			ReceiptsRoot:  executionPayload.ReceiptsRoot,
			LogsBloom:     executionPayload.LogsBloom,
			PrevRandao:    executionPayload.PrevRandao,
			BlockNumber:   executionPayload.BlockNumber,
			GasLimit:      executionPayload.GasLimit,
			GasUsed:       executionPayload.GasUsed,
			Timestamp:     executionPayload.Timestamp,
			ExtraData:     executionPayload.ExtraData,
			BaseFeePerGas: baseFeePerGasLE,
			Transactions:  executionPayload.Transactions,
		}

	case spec.DataVersionCapella.String():

		baseFeePerGas := [32]byte(executionPayload.BaseFeePerGas.PaddedBytes(32))
		baseFeePerGasLE := [32]byte{}
		for i := 0; i < len(baseFeePerGas); i++ {
			baseFeePerGasLE[i] = baseFeePerGas[len(baseFeePerGas)-1-i]
		}

		res.Capella = &capella.ExecutionPayload{
			ParentHash:    executionPayload.ParentHash,
			FeeRecipient:  executionPayload.FeeRecipient,
			StateRoot:     executionPayload.StateRoot,
			ReceiptsRoot:  executionPayload.ReceiptsRoot,
			LogsBloom:     executionPayload.LogsBloom,
			PrevRandao:    executionPayload.PrevRandao,
			BlockNumber:   executionPayload.BlockNumber,
			GasLimit:      executionPayload.GasLimit,
			GasUsed:       executionPayload.GasUsed,
			Timestamp:     executionPayload.Timestamp,
			ExtraData:     executionPayload.ExtraData,
			BaseFeePerGas: baseFeePerGasLE,
			BlockHash:     executionPayload.BlockHash,
			Transactions:  executionPayload.Transactions,
			Withdrawals:   executionPayload.Withdrawals,
		}

	case spec.DataVersionDeneb.String():
		res.Deneb = &deneb.ExecutionPayload{
			ParentHash:    executionPayload.ParentHash,
			FeeRecipient:  executionPayload.FeeRecipient,
			StateRoot:     executionPayload.StateRoot,
			ReceiptsRoot:  executionPayload.ReceiptsRoot,
			LogsBloom:     executionPayload.LogsBloom,
			PrevRandao:    executionPayload.PrevRandao,
			BlockNumber:   executionPayload.BlockNumber,
			GasLimit:      executionPayload.GasLimit,
			GasUsed:       executionPayload.GasUsed,
			Timestamp:     executionPayload.Timestamp,
			ExtraData:     executionPayload.ExtraData,
			BaseFeePerGas: executionPayload.BaseFeePerGas,
			BlockHash:     executionPayload.BlockHash,
			Transactions:  executionPayload.Transactions,
			Withdrawals:   executionPayload.Withdrawals,
			BlobGasUsed:   executionPayload.BlobGasUsed,
			ExcessBlobGas: executionPayload.ExcessBlobGas,
		}

	default:
		return res, errors.New("unknown fork version")

	}

	return res, nil
}

func ConstructExecutionPayloadV2(
	forkVersion string,
	executionPayload BaseExecutionPayload,
) (VersionedExecutionPayloadV2, error) {

	res := VersionedExecutionPayloadV2{}

	v1ExecutionPayload, err := ConstructExecutionPayload(forkVersion, executionPayload)
	if err != nil {
		return res, err
	}

	// Prior step would have already checked if the forkVersion is supported

	// So can now check for if deneb+ to add additional data
	v1ForkVersion, err := v1ExecutionPayload.VersionNumber()
	if err != nil {
		return res, err
	}
	if v1ForkVersion >= uint64(spec.DataVersionDeneb) {
		res.Deneb = &denebApi.ExecutionPayloadAndBlobsBundle{
			ExecutionPayload: v1ExecutionPayload.Deneb,
			BlobsBundle:      executionPayload.BlobsBundle,
		}
	} else {
		// Versions lower that deneb are forward compatible in conversion
		res, err = v1ExecutionPayload.ToVersionedExecutionPayloadV2()
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (v *VersionedExecutionPayload) ToBaseExecutionPayload() (BaseExecutionPayload, error) {
	res := BaseExecutionPayload{}

	switch {
	case v.Deneb != nil:
		res.ParentHash = v.Deneb.ParentHash
		res.FeeRecipient = v.Deneb.FeeRecipient
		res.StateRoot = v.Deneb.StateRoot
		res.ReceiptsRoot = v.Deneb.ReceiptsRoot
		res.LogsBloom = v.Deneb.LogsBloom
		res.PrevRandao = v.Deneb.PrevRandao
		res.BlockNumber = v.Deneb.BlockNumber
		res.GasLimit = v.Deneb.GasLimit
		res.GasUsed = v.Deneb.GasUsed
		res.Timestamp = v.Deneb.Timestamp
		res.ExtraData = v.Deneb.ExtraData
		res.BaseFeePerGas = v.Deneb.BaseFeePerGas
		res.BlockHash = v.Deneb.BlockHash
		res.Transactions = v.Deneb.Transactions
		res.Withdrawals = v.Deneb.Withdrawals
		res.BlobGasUsed = v.Deneb.BlobGasUsed
		res.ExcessBlobGas = v.Deneb.ExcessBlobGas

	case v.Capella != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(v.Capella.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = v.Capella.BaseFeePerGas[len(v.Capella.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}
		res.ParentHash = v.Capella.ParentHash
		res.FeeRecipient = v.Capella.FeeRecipient
		res.StateRoot = v.Capella.StateRoot
		res.ReceiptsRoot = v.Capella.ReceiptsRoot
		res.LogsBloom = v.Capella.LogsBloom
		res.PrevRandao = v.Capella.PrevRandao
		res.BlockNumber = v.Capella.BlockNumber
		res.GasLimit = v.Capella.GasLimit
		res.GasUsed = v.Capella.GasUsed
		res.Timestamp = v.Capella.Timestamp
		res.ExtraData = v.Capella.ExtraData
		res.BaseFeePerGas = baseFeePerGas
		res.BlockHash = v.Capella.BlockHash
		res.Transactions = v.Capella.Transactions
		res.Withdrawals = v.Capella.Withdrawals

	case v.Bellatrix != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBE := [32]byte{}
		for i := 0; i < len(v.Bellatrix.BaseFeePerGas); i++ {
			baseFeePerGasBE[i] = v.Bellatrix.BaseFeePerGas[len(v.Bellatrix.BaseFeePerGas)-1-i]
		}
		baseFeePerGasBigInt.SetBytes(baseFeePerGasBE[:])
		baseFeePerGas, overflow := uint256.FromBig(baseFeePerGasBigInt)
		if overflow {
			return res, errors.New("baseFeePerGas overflow")
		}
		res.ParentHash = v.Bellatrix.ParentHash
		res.FeeRecipient = v.Bellatrix.FeeRecipient
		res.StateRoot = v.Bellatrix.StateRoot
		res.ReceiptsRoot = v.Bellatrix.ReceiptsRoot
		res.LogsBloom = v.Bellatrix.LogsBloom
		res.PrevRandao = v.Bellatrix.PrevRandao
		res.BlockNumber = v.Bellatrix.BlockNumber
		res.GasLimit = v.Bellatrix.GasLimit
		res.GasUsed = v.Bellatrix.GasUsed
		res.Timestamp = v.Bellatrix.Timestamp
		res.ExtraData = v.Bellatrix.ExtraData
		res.BaseFeePerGas = baseFeePerGas
		res.BlockHash = v.Bellatrix.BlockHash
		res.Transactions = v.Bellatrix.Transactions
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil
}

func (v *VersionedExecutionPayload) ToVersionedExecutionPayloadV2() (VersionedExecutionPayloadV2, error) {
	// v1 is not forward compatible with v2 for deneb+
	// For deneb converting forward would result in missing blobs data
	res := VersionedExecutionPayloadV2{}

	switch {
	case v.Deneb != nil:
		res.Deneb = &denebApi.ExecutionPayloadAndBlobsBundle{
			ExecutionPayload: v.Deneb,
			BlobsBundle:      &denebApi.BlobsBundle{}, // Cannot obtain blobs data from v1
		}
	case v.Capella != nil:
		res.Capella = v.Capella
	case v.Bellatrix != nil:
		res.Bellatrix = v.Bellatrix
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil

}

func (v *VersionedExecutionPayloadV2) ToVersionedExecutionPayload() (VersionedExecutionPayload, error) {
	// v2 is backward compatible with all v1 versions
	res := VersionedExecutionPayload{}

	switch {
	case v.Deneb != nil:
		res.Deneb = v.Deneb.ExecutionPayload
	case v.Capella != nil:
		res.Capella = v.Capella
	case v.Bellatrix != nil:
		res.Bellatrix = v.Bellatrix
	default:
		return res, errors.New("unsupported fork version")
	}

	return res, nil

}

func (v *VersionedExecutionPayloadV2) ToBaseExecutionPayload() (BaseExecutionPayload, error) {
	res := BaseExecutionPayload{}

	v1ExecutionPayload, err := v.ToVersionedExecutionPayload()
	if err != nil {
		return res, err
	}

	res, err = v1ExecutionPayload.ToBaseExecutionPayload()
	if err != nil {
		return res, err
	}

	// Additional logic for after v1 execution payloads
	vNum, err := v.VersionNumber()
	if err != nil {
		return res, err
	}
	if vNum >= uint64(spec.DataVersionDeneb) {
		res.BlobsBundle = v.Deneb.BlobsBundle
	}

	return res, nil
}

// Converts the VersionedExecutionPayload to a VersionedExecutionPayloadHeader
func (v *VersionedExecutionPayload) ToVersionedExecutionPayloadHeader() (VersionedExecutionPayloadHeader, error) {
	res := VersionedExecutionPayloadHeader{}

	baseExecutionPayload, err := v.ToBaseExecutionPayload()
	if err != nil {
		return res, err
	}

	transactionsRoot, err := ComputeTransactionsRoot(baseExecutionPayload.Transactions)
	if err != nil {
		return res, err
	}

	withdrawalsRoot, err := ComputeWithdrawalsRoot(baseExecutionPayload.Withdrawals)
	if err != nil {
		return res, err
	}

	baseExecutionPayloadHeader := BaseExecutionPayloadHeader{
		ParentHash:       baseExecutionPayload.ParentHash,
		FeeRecipient:     baseExecutionPayload.FeeRecipient,
		StateRoot:        baseExecutionPayload.StateRoot,
		ReceiptsRoot:     baseExecutionPayload.ReceiptsRoot,
		LogsBloom:        baseExecutionPayload.LogsBloom,
		PrevRandao:       baseExecutionPayload.PrevRandao,
		BlockNumber:      baseExecutionPayload.BlockNumber,
		GasLimit:         baseExecutionPayload.GasLimit,
		GasUsed:          baseExecutionPayload.GasUsed,
		Timestamp:        baseExecutionPayload.Timestamp,
		ExtraData:        baseExecutionPayload.ExtraData,
		BaseFeePerGas:    baseExecutionPayload.BaseFeePerGas,
		BlockHash:        baseExecutionPayload.BlockHash,
		TransactionsRoot: transactionsRoot,
		WithdrawalsRoot:  withdrawalsRoot,
		BlobGasUsed:      baseExecutionPayload.BlobGasUsed,
		ExcessBlobGas:    baseExecutionPayload.ExcessBlobGas,
	}

	var forkVersion string
	switch {
	case v.Deneb != nil:
		forkVersion = spec.DataVersionDeneb.String()
	case v.Capella != nil:
		forkVersion = spec.DataVersionCapella.String()
	case v.Bellatrix != nil:
		forkVersion = spec.DataVersionBellatrix.String()
	default:
		return res, errors.New("unsupported fork version")
	}

	res, err = ConstructExecutionPayloadHeader(forkVersion, baseExecutionPayloadHeader)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *VersionedExecutionPayloadV2) ToVersionedExecutionPayloadHeader() (VersionedExecutionPayloadHeader, error) {
	res := VersionedExecutionPayloadHeader{}

	v1ExecutionPayload, err := v.ToVersionedExecutionPayload()
	if err != nil {
		return res, err
	}

	res, err = v1ExecutionPayload.ToVersionedExecutionPayloadHeader()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (v *VersionedExecutionPayload) Version() (string, error) {
	switch {
	case v.Bellatrix != nil:
		return spec.DataVersionBellatrix.String(), nil
	case v.Capella != nil:
		return spec.DataVersionCapella.String(), nil
	case v.Deneb != nil:
		return spec.DataVersionDeneb.String(), nil
	default:
		return "", errors.New("no fork version set")
	}
}

func (v *VersionedExecutionPayloadV2) Version() (string, error) {
	switch {
	case v.Bellatrix != nil:
		return spec.DataVersionBellatrix.String(), nil
	case v.Capella != nil:
		return spec.DataVersionCapella.String(), nil
	case v.Deneb != nil:
		return spec.DataVersionDeneb.String(), nil
	default:
		return "", errors.New("no fork version set")
	}
}

func (v *VersionedExecutionPayload) VersionNumber() (uint64, error) {
	switch {
	case v.Bellatrix != nil:
		return uint64(spec.DataVersionBellatrix), nil
	case v.Capella != nil:
		return uint64(spec.DataVersionCapella), nil
	case v.Deneb != nil:
		return uint64(spec.DataVersionDeneb), nil
	default:
		return 0, errors.New("no fork version set")
	}
}

func (v *VersionedExecutionPayloadV2) VersionNumber() (uint64, error) {
	switch {
	case v.Bellatrix != nil:
		return uint64(spec.DataVersionBellatrix), nil
	case v.Capella != nil:
		return uint64(spec.DataVersionCapella), nil
	case v.Deneb != nil:
		return uint64(spec.DataVersionDeneb), nil
	default:
		return 0, errors.New("no fork version set")
	}
}

func (v *VersionedExecutionPayload) WithVersionNumber() (VersionedExecutionPayloadWithVersionNumber, error) {
	res := VersionedExecutionPayloadWithVersionNumber{}

	versionNumber, err := v.VersionNumber()
	if err != nil {
		return res, err
	}

	res.VersionNumber = versionNumber
	res.VersionedExecutionPayload = v

	return res, nil

}

func (v *VersionedExecutionPayloadV2) WithVersionNumber() (VersionedExecutionPayloadV2WithVersionNumber, error) {
	res := VersionedExecutionPayloadV2WithVersionNumber{}

	versionNumber, err := v.VersionNumber()
	if err != nil {
		return res, err
	}

	res.VersionNumber = versionNumber
	res.VersionedExecutionPayload = v

	return res, nil

}

func (v *VersionedExecutionPayload) WithVersionName() (VersionedExecutionPayloadWithVersionName, error) {
	res := VersionedExecutionPayloadWithVersionName{}

	versionName, err := v.Version()
	if err != nil {
		return res, err
	}

	res.VersionName = versionName
	res.VersionedExecutionPayload = v

	return res, nil

}

func (v *VersionedExecutionPayloadV2) WithVersionName() (VersionedExecutionPayloadV2WithVersionName, error) {
	res := VersionedExecutionPayloadV2WithVersionName{}

	versionName, err := v.Version()
	if err != nil {
		return res, err
	}

	res.VersionName = versionName
	res.VersionedExecutionPayload = v

	return res, nil

}
