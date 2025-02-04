package request

type RequestShare struct {
	AppID           string      `json:"appid,omitempty"`
	TransactionID   string      `json:"transaction_id,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutOrderNO      string      `json:"out_order_no,omitempty"`
	Receivers       []*Receiver `json:"receivers,omitempty"`
	UnfreezeUnSplit bool        `json:"unfreeze_unsplit,omitempty"`
}

type Receiver struct {
	Type        string `json:"type"`
	Account     string `json:"account"`
	Name        string `json:"name,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}
