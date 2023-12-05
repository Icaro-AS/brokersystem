package dto

type TradeInput struct {
	OrderId       string  `json:"order_id"`
	InvestorId    string  `json:"investor_id"`
	AssetId       string  `json:"asset_id"`
	CurrentShares int     `jsob:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderId           string               `json:"order_id"`
	InvestorId        string               `json:"investor_id"`
	AssetId           string               `json:"asset_id"`
	OrderType         string               `json:"order_type"`
	Status            string               `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            int                  `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transaction_output"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerId       string  `json:"buyer_id"`
	SellerId      string  `json:"seller_id"`
	AssetId       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
