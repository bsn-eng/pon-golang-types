package common

import (
	"errors"
	"math/big"

	bellatrix "github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	deneb "github.com/attestantio/go-eth2-client/spec/deneb"
	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	uint256 "github.com/holiman/uint256"
)

// Always ensure that the execution payload header contains all
// the fields from all the fork versions, as they can be used or omitted
type BaseExecutionPayloadHeader struct {
	ParentHash       phase0.Hash32              `ssz-size:"32"`
	FeeRecipient     bellatrix.ExecutionAddress `ssz-size:"20"`
	StateRoot        phase0.Root                `ssz-size:"32"`
	ReceiptsRoot     phase0.Root                `ssz-size:"32"`
	LogsBloom        [256]byte                  `ssz-size:"256"`
	PrevRandao       [32]byte                   `ssz-size:"32"`
	BlockNumber      uint64
	GasLimit         uint64
	GasUsed          uint64
	Timestamp        uint64
	ExtraData        []byte        `ssz-max:"32"`
	BaseFeePerGas    *uint256.Int  `ssz-size:"32"`
	BlockHash        phase0.Hash32 `ssz-size:"32"`
	TransactionsRoot phase0.Root   `ssz-size:"32"`
	WithdrawalsRoot  phase0.Root   `ssz-size:"32"`
	BlobGasUsed      uint64
	ExcessBlobGas    uint64
}



func ConstructExecutionPayloadHeader(
	forkVersion string,
	executionPayloadHeader BaseExecutionPayloadHeader,
) (VersionedExecutionPayloadHeader, error) {

	res := VersionedExecutionPayloadHeader{}

	switch forkVersion {
	case "bellatrix":
		res.Bellatrix = &bellatrix.ExecutionPayloadHeader{
			ParentHash: executionPayloadHeader.ParentHash,
			FeeRecipient: executionPayloadHeader.FeeRecipient,
			StateRoot: executionPayloadHeader.StateRoot,
			ReceiptsRoot: executionPayloadHeader.ReceiptsRoot,
			LogsBloom: executionPayloadHeader.LogsBloom,
			PrevRandao: executionPayloadHeader.PrevRandao,
			BlockNumber: executionPayloadHeader.BlockNumber,
			GasLimit: executionPayloadHeader.GasLimit,
			GasUsed: executionPayloadHeader.GasUsed,
			Timestamp: executionPayloadHeader.Timestamp,
			ExtraData: executionPayloadHeader.ExtraData,
			BaseFeePerGas: executionPayloadHeader.BaseFeePerGas.Bytes32(),
			BlockHash: executionPayloadHeader.BlockHash,
			TransactionsRoot: executionPayloadHeader.TransactionsRoot,
		}
	
	case "capella":
		res.Capella = &capella.ExecutionPayloadHeader{
			ParentHash: executionPayloadHeader.ParentHash,
			FeeRecipient: executionPayloadHeader.FeeRecipient,
			StateRoot: executionPayloadHeader.StateRoot,
			ReceiptsRoot: executionPayloadHeader.ReceiptsRoot,
			LogsBloom: executionPayloadHeader.LogsBloom,
			PrevRandao: executionPayloadHeader.PrevRandao,
			BlockNumber: executionPayloadHeader.BlockNumber,
			GasLimit: executionPayloadHeader.GasLimit,
			GasUsed: executionPayloadHeader.GasUsed,
			Timestamp: executionPayloadHeader.Timestamp,
			ExtraData: executionPayloadHeader.ExtraData,
			BaseFeePerGas: executionPayloadHeader.BaseFeePerGas.Bytes32(),
			BlockHash: executionPayloadHeader.BlockHash,
			TransactionsRoot: executionPayloadHeader.TransactionsRoot,
			WithdrawalsRoot: executionPayloadHeader.WithdrawalsRoot,
		}
	
	case "deneb":
		res.Deneb = &deneb.ExecutionPayloadHeader{
			ParentHash: executionPayloadHeader.ParentHash,
			FeeRecipient: executionPayloadHeader.FeeRecipient,
			StateRoot: executionPayloadHeader.StateRoot,
			ReceiptsRoot: executionPayloadHeader.ReceiptsRoot,
			LogsBloom: executionPayloadHeader.LogsBloom,
			PrevRandao: executionPayloadHeader.PrevRandao,
			BlockNumber: executionPayloadHeader.BlockNumber,
			GasLimit: executionPayloadHeader.GasLimit,
			GasUsed: executionPayloadHeader.GasUsed,
			Timestamp: executionPayloadHeader.Timestamp,
			ExtraData: executionPayloadHeader.ExtraData,
			BaseFeePerGas: executionPayloadHeader.BaseFeePerGas,
			BlockHash: executionPayloadHeader.BlockHash,
			TransactionsRoot: executionPayloadHeader.TransactionsRoot,
			WithdrawalsRoot: executionPayloadHeader.WithdrawalsRoot,
			BlobGasUsed: executionPayloadHeader.BlobGasUsed,
			ExcessBlobGas: executionPayloadHeader.ExcessBlobGas,
		}

	default:
		return res, errors.New("unknown fork version")

	}

	return res, nil
}

func (v *VersionedExecutionPayloadHeader) ToBaseExecutionPayloadHeader() (BaseExecutionPayloadHeader, error) {
	res := BaseExecutionPayloadHeader{}

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
		res.TransactionsRoot = v.Deneb.TransactionsRoot
		res.WithdrawalsRoot = v.Deneb.WithdrawalsRoot
		res.BlobGasUsed = v.Deneb.BlobGasUsed
		res.ExcessBlobGas = v.Deneb.ExcessBlobGas

	case v.Capella != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBigInt.SetBytes(v.Capella.BaseFeePerGas[:])
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
		res.TransactionsRoot = v.Capella.TransactionsRoot
		res.WithdrawalsRoot = v.Capella.WithdrawalsRoot

	case v.Bellatrix != nil:
		baseFeePerGasBigInt := big.NewInt(0)
		baseFeePerGasBigInt.SetBytes(v.Bellatrix.BaseFeePerGas[:])
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
	default:
		return res, errors.New("unknown fork version")
	}

	return res, nil
}


// Converts the VersionedExecutionPayloadHeader to a VersionedExecutionPayload without the transactions and withdrawals as those cannot be derived
func (v *VersionedExecutionPayloadHeader) ToVersionedExecutionPayload() (VersionedExecutionPayload, error) {
	res := VersionedExecutionPayload{}

	baseExecutionPayloadHeader, err := v.ToBaseExecutionPayloadHeader()
	if err != nil {
		return res, err
	}

	baseExecutionPayload := BaseExecutionPayload{
		ParentHash: baseExecutionPayloadHeader.ParentHash,
		FeeRecipient: baseExecutionPayloadHeader.FeeRecipient,
		StateRoot: baseExecutionPayloadHeader.StateRoot,
		ReceiptsRoot: baseExecutionPayloadHeader.ReceiptsRoot,
		LogsBloom: baseExecutionPayloadHeader.LogsBloom,
		PrevRandao: baseExecutionPayloadHeader.PrevRandao,
		BlockNumber: baseExecutionPayloadHeader.BlockNumber,
		GasLimit: baseExecutionPayloadHeader.GasLimit,
		GasUsed: baseExecutionPayloadHeader.GasUsed,
		Timestamp: baseExecutionPayloadHeader.Timestamp,
		ExtraData: baseExecutionPayloadHeader.ExtraData,
		BaseFeePerGas: baseExecutionPayloadHeader.BaseFeePerGas,
		BlockHash: baseExecutionPayloadHeader.BlockHash,
		// Transactions:
		// Withdrawals:
		BlobGasUsed: baseExecutionPayloadHeader.BlobGasUsed,
		ExcessBlobGas: baseExecutionPayloadHeader.ExcessBlobGas,
	}

	var forkVersion string
	switch {
	case v.Deneb != nil:
		forkVersion = "deneb"
	case v.Capella != nil:
		forkVersion = "capella"
	case v.Bellatrix != nil:
		forkVersion = "bellatrix"
	default:
		return res, errors.New("unknown fork version")
	}

	res, err = ConstructExecutionPayload(forkVersion, baseExecutionPayload)
	if err != nil {
		return res, err
	}

	return res, nil
}