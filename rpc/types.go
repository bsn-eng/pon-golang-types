package rpc


// RPC Params

type JSONrpcPrivateRawTxs []string

type JSONrpcPrivateTx struct {
	Tx string `json:"tx"`
}

type JSONrpcPrivateTxs []JSONrpcPrivateTx

type JSONrpcPrivateTxHash struct {
	TxHash string `json:"txHash"`
}

type JSONrpcPrivateTxHashes []JSONrpcPrivateTxHash

type JSONrpcBundle struct {
	ID string `json:"id,omitempty"`  // ID is the bundle ID
	Txs []string `json:"txs,omitempty"`  // Hex-encoded transaction bytes
	BlockNumber uint64 `json:"blockNumber,string,omitempty"`
	MinTimestamp uint64 `json:"minTimestamp,string,omitempty"`
	MaxTimestamp uint64 `json:"maxTimestamp,string,omitempty"`
	RevertingTxHashes []string `json:"revertingTxHashes,omitempty"`
}

type JSONrpcBundleHash struct {
	BundleHash string `json:"bundleHash"`
}

type JSONrpcBundleHashes []JSONrpcBundleHash