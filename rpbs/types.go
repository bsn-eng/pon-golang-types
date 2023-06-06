package rpbs

type RPBSCommitMessage struct {
	BuilderWalletAddress string `json:"builderWalletAddress"`
	Slot                 uint64 `json:"slot"`
	Amount               uint64 `json:"amount"`
	PayoutTxBytes        string `json:"payoutTxBytes"`
	TxBytes              string `json:"txBytes"`
}

type EncodedRPBSSignature struct {
	Z1Hat string `json:"z1_hat"`
	C1Hat string `json:"c1_hat"`
	S1Hat string `json:"s1_hat"`
	C2Hat string `json:"c2_hat"`
	S2Hat string `json:"s2_hat"`
	M1Hat string `json:"m1_hat"`
}