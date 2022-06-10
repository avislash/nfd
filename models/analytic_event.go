package models

type AnalyticEvent struct {
	Block    int      `json:"block"`
	Buyer    string   `json:"buyer"`
	Category Category `json:"category"`
	Event    Event    `json:"event"`
	GroupID  string   `json:"groupID"`
	MetaTags []string `json:"metaTags"`
	Name     string   `json:"name"`
	Note     string   `json:"note"`
	SaleType string   `json:"saleType"`
	Seller   string   `json:"seller"`
}
