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

type SaleType string

const (
	AUCTION    SaleType = "auction"
	BUY_IT_NOW SaleType = "buyItNow"
)

type State string

const (
	AVAILABLE State = "available"
	MINTING   State = "minting"
	RESERVED  State = "reserved"
	FOR_SALE  State = "forSale"
	OWNED     State = "owned"
)
