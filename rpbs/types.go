package rpbs

type RPBSCommitMessage struct {
	BuilderWalletAddress string `json:"builderWalletAddress"`
	Slot                 uint64 `json:"slot"`
	Amount               uint64 `json:"amount"`
	PayoutTxBytes        string `json:"payoutTxBytes"`
	TxBytes              string `json:"txBytes"`
}

type EncodedRPBSSignature struct {
	Z1Hat string `json:"z1Hat"`
	C1Hat string `json:"c1Hat"`
	S1Hat string `json:"s1Hat"`
	C2Hat string `json:"c2Hat"`
	S2Hat string `json:"s2Hat"`
	M1Hat string `json:"m1Hat"`
}