package response

type TransactionResponse struct {
	Hash      string `json:"hash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Gas       uint64 `json:"gas"`
	GasPrice  string `json:"gasPrice"`
	Nonce     uint64 `json:"nonce"`
	IsPending bool   `json:"isPending"`
}

type BlockResponse struct {
	Number       string   `json:"number"`
	Hash         string   `json:"hash"`
	ParentHash   string   `json:"parentHash"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type ContractValueResponse struct {
	Value string `json:"value"`
}

type SetContractValueResponse struct {
	TransactionHash string `json:"transactionHash"`
}
