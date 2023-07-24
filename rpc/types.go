package rpc


// RPC Params

type JSONrpcPrivateRawTxs struct {
	Txs []string `json:"txs"`
}

type JSONrpcPrivateTx struct {
	Tx string `json:"tx"`
}

type JSONrpcPrivateTxs struct {
	Txs []JSONrpcPrivateTx `json:"txs"`
}

type JSONrpcPrivateTxHash struct {
	TxHash string `json:"txHash"`
}

type JSONrpcPrivateTxHashes struct {
	TxHashes []JSONrpcPrivateTxHash `json:"txHashes"`
}

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

type JSONrpcBundleHashes struct {
	BundleHashes []JSONrpcBundleHash `json:"bundleHashes"`
}