package RPBS

import relayTypes "github.com/bsn-eng/pon-golang-types/relay"

type RpbsCommitResponse map[string]string
type RpbsChallengeResponse map[string]string
type RpbsSolution map[string]string

type RpbsCommitMessage struct {
	BuilderWalletAddress *relayTypes.Address `json:"builderWalletAddress"`
	Slot                 uint64              `json:"slot"`
	Amount               uint64              `json:"amount"`
	TxBytes              string              `json:"txBytes"`
}

type RPBSChallenge struct {
	Commitment RpbsCommitResponse    `json:"commitment"`
	Challenge  RpbsChallengeResponse `json:"challenge"`
}

type ReporterSlot struct {
	SlotLower uint64 `json:"slot_lower"`
	SlotUpper uint64 `json:"slot_upper"`
}

type GetValidatorsResponseEntry struct {
	Slot   uint64 `json:"slot,string"`
	PubKey string `json:"pubKey"`
}
