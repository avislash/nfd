package client

import (
	"net/url"
	"strconv"
	"time"
)

type AnalyticsOptions struct {
	Name         string
	Buyer        string
	Seller       string
	Event        Event
	RequireBuyer bool
	Category     Category
	SaleType     SaleType
	State        State
	Length       Length
	Traits       []Trait
	MinPrice     int
	MaxPrice     int
	After        time.Time
	Limit        int
	Offset       int
	Sort         SortOption
	IfNoneMatch  string
}

type SuggestOptions struct {
	Limit int
	View  View
}

func (so *SuggestOptions) ToQueryParams() url.Values {
	queryParams := make(url.Values)
	if so.Limit != 0 {
		queryParams.Add("limit", strconv.Itoa(so.Limit))
	}

	if len(so.View) != 0 {
		switch so.View {
		case BRIEF_VIEW, FULL_VIEW:
			queryParams.Add("view", string(so.View))
		}
	}

	return queryParams
}

func (ao *AnalyticsOptions) ToQueryParams() url.Values {
	queryParams := make(url.Values)
	if len(ao.Name) != 0 {
		queryParams.Add("name", ao.Name)
	}

	if len(ao.Buyer) != 0 {
		queryParams.Add("buyer", ao.Buyer)
	}

	if len(ao.Seller) != 0 {
		queryParams.Add("seller", ao.Seller)
	}

	if len(ao.Event) != 0 {
		queryParams.Add("event", string(ao.Event))
	}

	if ao.RequireBuyer {
		queryParams.Add("requireBuyer", "true")
	}

	if len(ao.Category) != 0 {
		queryParams.Add("category", string(ao.Category))
	}

	if len(ao.SaleType) != 0 {
		queryParams.Add("saleType", string(ao.SaleType))
	}

	if len(ao.State) != 0 {
		queryParams.Add("state", string(ao.State))
	}

	if len(ao.Length) != 0 {
		queryParams.Add("length", string(ao.Length))
	}

	if len(ao.Traits) != 0 {
		for _, trait := range ao.Traits {
			queryParams.Add("traits", string(trait))
		}
	}

	if 0 != ao.MinPrice {
		queryParams.Add("minPrice", strconv.Itoa(ao.MinPrice))
	}
	if 0 != ao.MaxPrice {
		queryParams.Add("maxPrice", strconv.Itoa(ao.MaxPrice))
	}

	if !ao.After.IsZero() {
		queryParams.Add("afterTime", ao.After.Format(time.RFC3339))
	}

	if 0 != ao.Limit {
		queryParams.Add("limit", strconv.Itoa(ao.Limit))
	}

	if 0 != ao.Offset {
		queryParams.Add("offset", strconv.Itoa(ao.Offset))
	}

	if len(ao.Sort) != 0 {
		switch ao.Sort {
		case DESCENDING_TIME_SORT, ASCENDING_PRICE_SORT, DESECENDING_PRICE_SORT:
			queryParams.Add("sort", string(ao.Sort))
		}
	}

	return queryParams
}
