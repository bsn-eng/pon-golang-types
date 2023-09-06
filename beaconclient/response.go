package beaconclient

import (
	"github.com/bsn-eng/pon-golang-types/common"
)

type GetGenesisResponse struct {
	Data *GenesisData `json:"data"`
}

type GetWithdrawalsResponse struct {
	Data *Withdrawals `json:"data"`
}

type GetSyncStatusResponse struct {
	Data *SyncStatusData
}

type GetValidatorsResponse struct {
	Data []*ValidatorData `json:"data"`
}

type GetProposerDutiesResponse struct {
	Data []*ProposerDutyData `json:"data"`
}

type GetRandaoResponse struct {
	Data *RandaoData `json:"data"`
}

type GetBlockResponse struct {
	Version string `json:"version"`
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized bool `json:"finalized"`
	Data *common.VersionedSignedBeaconBlock `json:"data"`
}

type GetBlockHeaderResponse struct {
	Data *BlockHeaderData `json:"data"`
}