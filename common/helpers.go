package common

import (
	"errors"

	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	capella "github.com/attestantio/go-eth2-client/spec/capella"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	utilbellatrix "github.com/attestantio/go-eth2-client/util/bellatrix"
	utilcapella "github.com/attestantio/go-eth2-client/util/capella"
	"github.com/ethereum/go-ethereum/common"
)

var (
	EmptyWithdrawalMerkleRoot = "0x792930bbd5baac43bcc798ee49aa8185ef76bb3b44ba62b91d86ae569e4bb535"
)

func ComputeWithdrawalsRoot(w []*capella.Withdrawal) (phase0.Root, error) {
	if len(w) == 0 {
		emptyRoot := phase0.Root{}
		emptyHash := common.HexToHash(EmptyWithdrawalMerkleRoot)
		copy(emptyRoot[:], emptyHash.Bytes()[:])
		return emptyRoot, nil
	}
	withdrawals := utilcapella.ExecutionPayloadWithdrawals{Withdrawals: w}
	return withdrawals.HashTreeRoot()
}

func ComputeTransactionsRoot(t []bellatrix.Transaction) (phase0.Root, error) {

	transactions := utilbellatrix.ExecutionPayloadTransactions{Transactions: t}
	return transactions.HashTreeRoot()
}