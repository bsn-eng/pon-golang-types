package ponpool

type Builder struct {
	BuilderPubkey string `db:"builder_pubkey" json:"id"`
	Status        string `db:"status" json:"status"`
	BalanceStaked string `db:"balance_staked" json:"balanceStaked"`
}

type BuilderStake struct {
	Data struct {
		GlobalValue struct {
			BalanceRequired string `json:"builderMinimumStake"`
		} `json:"globalValue"`
	} `json:"data"`
}

type BuilderInterface struct {
	Builder Builder
	Status  bool
}

type BuilderPool struct {
	Data struct {
		Builders []Builder `json:"builders"`
	} `json:"data"`
}

type Validator struct {
	ValidatorPubkey string `db:"validator_pubkey" json:"id"`
	Status          string `db:"status" json:"status"`
	ReportCount     string `db:"report_count" json:"reportCount"`
}

type ValidatorPool struct {
	Data struct {
		Validators []Validator `json:"proposers"`
	} `json:"data"`
}

type Reporter struct {
	ReporterPubkey string `db:"reporter_pubkey" json:"id"`
	Active         bool   `db:"active" json:"active"`
	ReportCount    string `db:"report_count" json:"numberOfReports"`
}

type ReporterPool struct {
	Data struct {
		Reporters []Reporter `json:"reporters"`
	} `json:"data"`
}
