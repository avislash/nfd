package client

import (
	"net/url"
	"strconv"
	"time"
)

type option interface {
	GetQueryParams() url.Values
	GetHeader() map[string]string
}

type ActivityOptions struct {
	Type        Type
	After       time.Time
	Limit       int
	Sort        SortOption
	IfNoneMatch string
}

func (ao *ActivityOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if ao == nil {
		return queryParams
	}

	if len(ao.Type) != 0 {
		queryParams.Add("type", string(ao.Type))
	}

	if !ao.After.IsZero() {
		queryParams.Add("afterTime", ao.After.Format(time.RFC3339))
	}

	if ao.Limit != 0 {
		queryParams.Add("limit", strconv.Itoa(ao.Limit))
	}

	if len(ao.Sort) != 0 && DESCENDING_TIME_SORT == ao.Sort {
		queryParams.Add("sort", string(ao.Sort))
	}

	return queryParams
}

func (ao *ActivityOptions) GetHeader() map[string]string {
	var header map[string]string

	if nil != ao && len(ao.IfNoneMatch) != 0 {
		header = map[string]string{
			"if-none-match": ao.IfNoneMatch,
		}
	}

	return header
}

type AddressOptions struct {
	Limit       int
	View        View
	IfNoneMatch string
}

func (ao *AddressOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil == ao {
		return queryParams
	}

	if 0 != ao.Limit {
		queryParams.Add("limit", strconv.Itoa(ao.Limit))
	}

	if len(ao.View) != 0 {
		queryParams.Add("view", string(ao.View))
	}

	return queryParams
}

func (ao *AddressOptions) GetHeader() map[string]string {
	var header map[string]string

	if nil == ao {
		return header
	}

	if len(ao.IfNoneMatch) != 0 {
		header = map[string]string{
			"if-none-match": ao.IfNoneMatch,
		}
	}

	return header
}

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

func (ao *AnalyticsOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil == ao {
		return queryParams
	}

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

func (ao *AnalyticsOptions) GetHeader() map[string]string {
	var header map[string]string
	if nil != ao && len(ao.IfNoneMatch) != 0 {
		header = map[string]string{
			"if-none-match": ao.IfNoneMatch,
		}
	}
	return header
}

type AuctionOptions struct {
	Name         string
	StartingSoon bool
}

func (ao *AuctionOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil == ao {
		return queryParams
	}

	if len(ao.Name) != 0 {
		queryParams.Add("name", ao.Name)
	}

	if ao.StartingSoon {
		queryParams.Add("startingSoon", "true")
	}

	return queryParams
}

func (ao *AuctionOptions) GetHeader() map[string]string {
	return nil
}

type BrowseOptions struct {
	Name        string
	Category    Category
	SaleType    SaleType
	State       State
	Length      Length
	Traits      []Trait
	MinPrice    int
	MaxPrice    int
	Limit       int
	Offset      int
	Sort        SortOption
	View        View
	IfNoneMatch string
}

func (bo *BrowseOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil == bo {
		return queryParams
	}

	if len(bo.Name) != 0 {
		queryParams.Add("name", bo.Name)
	}

	if len(bo.Category) != 0 {
		queryParams.Add("category", string(bo.Category))
	}

	if len(bo.SaleType) != 0 {
		queryParams.Add("saleType", string(bo.SaleType))
	}

	if len(bo.State) != 0 {
		queryParams.Add("state", string(bo.State))
	}

	if len(bo.Length) != 0 {
		queryParams.Add("length", string(bo.Length))
	}

	if len(bo.Traits) != 0 {
		for _, trait := range bo.Traits {
			queryParams.Add("trait", string(trait))
		}
	}

	if 0 != bo.MinPrice {
		queryParams.Add("minPrice", strconv.Itoa(bo.MinPrice))
	}

	if 0 != bo.MaxPrice {
		queryParams.Add("maxPrice", strconv.Itoa(bo.MaxPrice))
	}

	if 0 != bo.Limit {
		queryParams.Add("limit", strconv.Itoa(bo.Limit))
	}

	if 0 != bo.Offset {
		queryParams.Add("offset", strconv.Itoa(bo.Offset))
	}

	if len(bo.Sort) != 0 && DESCENDING_TIME_SORT != bo.Sort {
		queryParams.Add("sort", string(bo.Sort))
	}

	if len(bo.View) != 0 && THUMBNAIL_VIEW != bo.View {
		queryParams.Add("view", string(bo.View))
	}

	return queryParams
}

func (bo *BrowseOptions) GetHeader() map[string]string {
	var header map[string]string

	if nil == bo {
		return header
	}

	if len(bo.IfNoneMatch) != 0 {
		header = map[string]string{
			"if-none-match": bo.IfNoneMatch,
		}
	}
	return header
}

type NameOptions struct {
	View        View
	Poll        bool
	NoCache     bool
	IfNoneMatch string
}

func (no *NameOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil != no {
		return queryParams
	}

	if len(no.View) != 0 {
		queryParams.Add("view", string(no.View))
	}

	if false != no.Poll {
		queryParams.Add("poll", "true")
	}

	if false != no.NoCache {
		queryParams.Add("nocache", "true")
	}

	return queryParams
}

func (no *NameOptions) GetHeader() map[string]string {
	var header map[string]string
	if nil != no && len(no.IfNoneMatch) != 0 {
		header = map[string]string{
			"if-none-match": no.IfNoneMatch,
		}
	}
	return header
}

type SuggestOptions struct {
	Limit int
	View  View
}

func (so *SuggestOptions) GetQueryParams() url.Values {
	queryParams := make(url.Values)

	if nil == so {
		return queryParams
	}

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

func (so *SuggestOptions) GetHeader() map[string]string {
	return nil
}
