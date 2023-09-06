package beaconclient

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

type GetBlockHeaderResponse struct {
	Data *BlockHeaderData `json:"data"`
}