package request

type TransactionsRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type ContractValueRequest struct {
	From  string `json:"from"`
	Value uint64 `json:"value"`
}
