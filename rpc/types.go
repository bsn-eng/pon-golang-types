package rpc


// RPC Params

type JSONrpcPrivateRawTxs []string

type JSONrpcPrivateTx struct {
	Tx string `json:"tx"`
}

type JSONrpcPrivateTxs []JSONrpcPrivateTx

type JSONrpcBundle struct {
	Txs []string `json:"txs"`  // Hex-encoded transaction bytes
	BlockNumber uint64 `json:"blockNumber,string"`
	MinTimestamp uint64 `json:"minTimestamp,string"`
	MaxTimestamp uint64 `json:"maxTimestamp,string"`
	RevertingTxHashes []string `json:"revertingTxHashes"`
	BundleCreatorPubkey string `json:"bundleCreatorPubkey"`
	BundleCreatorSignature string `json:"bundleCreatorSignature"`
}

type JSONrpcBundleHash struct {
	BundleHash string `json:"bundleHash"`
}

type JSONrpcBundleHashes []JSONrpcBundleHash