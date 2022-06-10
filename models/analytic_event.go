package models

type Category string

const (
	CURATED Category = "curated"
	PREMIUM Category = "premium"
	COMMON  Category = "common"
)

type Event string

const (
	MINTED           Event = "minted"
	OFFERED_FOR_SALE Event = "offeredForSale"
	CANCELED_SALE    Event = "canceledSale"
	SOLD             Event = "sold"
	POSTED_OFFER     Event = "postedOffer"
)

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
