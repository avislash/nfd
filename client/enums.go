package client

import "github.com/avislash/nfd/models"

type Event = models.Event
type Category = models.Category

type View string

const (
	BRIEF_VIEW     View = "brief"
	TINY_VIEW      View = "tiny"
	THUMBNAIL_VIEW View = "thumbnail"
	FULL_VIEW      View = "full"
)

type Type string

const (
	CHANGES Type = "changes"
)

type SortOption string

const (
	DESCENDING_TIME_SORT         SortOption = "timeDesc"
	ASCENDING_PRICE_SORT         SortOption = "priceAsc"
	ASCENDING_SALE_TYPE_SORT     SortOption = "saleTypeAsc"
	DESECENDING_PRICE_SORT       SortOption = "priceDesc"
	DESCENDING_CREATED_SORT      SortOption = "createdDesc"
	DESCENDING_SOLD_SORT         SortOption = "soldDesc"
	DESCENDING_HIGHEST_SALE_SORT SortOption = "highestSaleDesc"
)

type SaleType string

const (
	AUCTION    SaleType = "auction"
	BUY_IT_NOW SaleType = "buyItNow"
)

type State string

const (
	MINTING  State = "minting"
	FOR_SALE State = "forSale"
	OWNED    State = "owned"
)

type Length string

const (
	ONE_LETTER       Length = "1_letters"
	TWO_LETTERS      Length = "2_letters"
	THREE_LETTERS    Length = "3_letters"
	FOUR_LETTERS     Length = "4_letters"
	FIVE_LETTERS     Length = "5_letters"
	SIX_LETTERS      Length = "6_letters"
	SEVEN_LETTERS    Length = "7_letters"
	EIGHT_LETTERS    Length = "8_letters"
	NINE_LETTERS     Length = "9_letters"
	TEN_PLUS_LETTERS Length = "10+_letters"
)

type Trait string

const (
	EMOJI Trait = "emoji"
)
