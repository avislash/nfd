package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/avislash/nfd/models"
)

const (
	baseURL string = "https://api.nf.domains/"
)

type Client struct {
	*http.Client
	baseURL string
}

func NewClient() *Client {
	return &Client{
		&http.Client{},
		baseURL,
	}
}

func (c *Client) Version() (models.Version, error) {
	var version models.Version
	resp, err := c.get(baseURL+"info/version", nil)
	if err != nil {
		return version, fmt.Errorf("Error fetching version: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&version)
	return version, err
}

func (c *Client) Totals() (models.Totals, error) {
	var totals models.Totals

	resp, err := c.get(baseURL+"nfd/totals", nil)
	if err != nil {
		return totals, fmt.Errorf("Error fetching totals: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&totals)
	return totals, err
}

//TODO: Find out what if-none-match is
func (c *Client) Name(name string, options *NameOptions) (models.Record, error) {
	var record models.Record
	var header map[string]string
	query, _ := url.Parse(baseURL + "nfd/" + name)
	queryParams := query.Query()
	if nil != options {
		if len(options.View) != 0 {
			queryParams.Add("view", string(options.View))
		}

		if false != options.Poll {
			queryParams.Add("poll", "true")
		}

		if false != options.NoCache {
			queryParams.Add("nocache", "true")
		}
		if len(options.IfNoneMatch) != 0 {
			header = map[string]string{
				"if-none-match": options.IfNoneMatch,
			}
		}
	}
	query.RawQuery = queryParams.Encode()

	resp, err := c.get(query.String(), header)
	if err != nil {
		return record, fmt.Errorf("Unable to fetch NFD Record for %s: %w", name, err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&record)

	return record, err
}

func (c *Client) Activity(names []string, options *ActivityOptions) ([]models.ActivityRecord, error) {
	var activityRecords []models.ActivityRecord
	var header map[string]string
	query, _ := url.Parse(baseURL + "nfd/activity")

	queryParams := query.Query()
	for _, name := range names {
		queryParams.Add("name", name)
	}

	if nil != options {
		if len(options.Type) != 0 {
			queryParams.Add("type", string(options.Type))
		}

		if !options.After.IsZero() {
			queryParams.Add("afterTime", options.After.Format(time.RFC3339))
		}

		if 0 != options.Limit {
			queryParams.Add("limit", strconv.Itoa(options.Limit))
		}

		if len(options.Sort) != 0 && DESCENDING_TIME_SORT == options.Sort {
			queryParams.Add("sort", string(options.Sort))
		}

		if len(options.IfNoneMatch) != 0 {
			header = map[string]string{
				"if-none-match": options.IfNoneMatch,
			}
		}
	}
	query.RawQuery = queryParams.Encode()

	resp, err := c.get(query.String(), header)
	if err != nil {
		return activityRecords, fmt.Errorf("Unable to get Activity Records for %s: %w", names, err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&activityRecords)

	return activityRecords, err

}

//TODO: Add tags to all options and then just marshal the options into a map and walk through the map to set the values
type AddressOptions struct {
	Limit       int
	View        View
	IfNoneMatch string
}

func (c *Client) Address(addresses []string, options *AddressOptions) ([]models.Record, error) {
	var records []models.Record
	var header map[string]string
	query, _ := url.Parse(baseURL + "nfd/address")

	queryParams := query.Query()
	for _, address := range addresses {
		queryParams.Add("address", address)
	}

	if nil != options {
		if 0 != options.Limit {
			queryParams.Add("limit", strconv.Itoa(options.Limit))
		}

		if len(options.View) != 0 {
			queryParams.Add("view", string(options.View))
		}

		if len(options.IfNoneMatch) != 0 {
			header = map[string]string{
				"if-none-match": options.IfNoneMatch,
			}
		}
	}
	query.RawQuery = queryParams.Encode()

	resp, err := c.get(query.String(), header)
	if err != nil {
		return records, fmt.Errorf("Unable to get Records for %s: %w", addresses, err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&records)

	return records, err
}

func (c *Client) Analytics(options *AnalyticsOptions) (models.AnalyticRecords, error) {
	var records models.AnalyticRecords
	var header map[string]string
	query, _ := url.Parse(baseURL + "nfd/analytics")

	if options != nil {
		query.RawQuery = options.ToQueryParams().Encode()
		if len(options.IfNoneMatch) != 0 {
			header = map[string]string{
				"if-none-match": options.IfNoneMatch,
			}
		}
	}

	resp, err := c.get(query.String(), header)
	if err != nil {
		return records, fmt.Errorf("Unable to get Analytics Records: %w", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&records)

	return records, err
}

func (c *Client) Auction(options *AuctionOptions) ([]models.AuctionAndPrice, error) {
	var openAuctions []models.AuctionAndPrice
	query, _ := url.Parse(baseURL + "nfd/auction")

	queryParams := query.Query()
	if options != nil {
		if len(options.Name) != 0 {
			queryParams.Add("name", options.Name)
		}

		if options.StartingSoon {
			queryParams.Add("startingSoon", "true")
		}
	}
	query.RawQuery = queryParams.Encode()
	resp, err := c.get(query.String(), nil)
	if err != nil {
		return openAuctions, fmt.Errorf("Unable to get Auction and Price: %w", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&openAuctions)

	return openAuctions, err
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

func (c *Client) Browse(options *BrowseOptions) ([]models.Record, error) {
	var records []models.Record
	var header map[string]string
	query, _ := url.Parse(baseURL + "nfd/browse")

	queryParams := query.Query()
	if options != nil {
		if len(options.Name) != 0 {
			queryParams.Add("name", options.Name)
		}

		if len(options.Category) != 0 {
			queryParams.Add("category", string(options.Category))
		}

		if len(options.SaleType) != 0 {
			queryParams.Add("saleType", string(options.SaleType))
		}

		if len(options.State) != 0 {
			queryParams.Add("state", string(options.State))
		}

		if len(options.Length) != 0 {
			queryParams.Add("length", string(options.Length))
		}

		if len(options.Traits) != 0 {
			for _, trait := range options.Traits {
				queryParams.Add("trait", string(trait))
			}
		}

		if 0 != options.MinPrice {
			queryParams.Add("minPrice", strconv.Itoa(options.MinPrice))
		}

		if 0 != options.MaxPrice {
			queryParams.Add("maxPrice", strconv.Itoa(options.MaxPrice))
		}

		if 0 != options.Limit {
			queryParams.Add("limit", strconv.Itoa(options.Limit))
		}

		if 0 != options.Offset {
			queryParams.Add("offset", strconv.Itoa(options.Offset))
		}

		if len(options.Sort) != 0 && DESCENDING_TIME_SORT != options.Sort {
			queryParams.Add("sort", string(options.Sort))
		}

		if len(options.View) != 0 && THUMBNAIL_VIEW != options.View {
			queryParams.Add("view", string(options.View))
		}

		if len(options.IfNoneMatch) != 0 {
			header = map[string]string{
				"if-none-match": options.IfNoneMatch,
			}
		}
	}
	query.RawQuery = queryParams.Encode()

	resp, err := c.get(query.String(), header)
	if err != nil {
		return records, fmt.Errorf("Unable to Browse NFD: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&records)

	return records, nil

}

func (c *Client) IsValid(appID int) (models.Valid, error) {
	var isValid models.Valid
	query, _ := url.Parse(baseURL + "nfd/isValid/" + strconv.Itoa(appID))

	resp, err := c.get(query.String(), nil)
	if err != nil {
		return isValid, fmt.Errorf("Unable check Validity for %d: %w", appID, err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&isValid)

	return isValid, err
}

func (c *Client) NameSig(name string) (string, error) {
	query, _ := url.Parse(baseURL + "nfd/nameSig/" + name)

	resp, err := c.get(query.String(), nil)
	if err != nil {
		return "", fmt.Errorf("Unable to Get Name Sig for %s: %w", name, err)
	}
	defer resp.Body.Close()

	nameSig, err := io.ReadAll(resp.Body)

	return string(nameSig[1 : len(nameSig)-2]), err //TODO: Fix hack for sanitzing string
}

func (c *Client) RevSig(address string) (string, error) {
	query, err := url.Parse(baseURL + "nfd/revAddressSig/" + address)
	if err != nil {
		return "", err
	}

	resp, err := c.get(query.String(), nil)
	if err != nil {
		return "", fmt.Errorf("Unable to Get Rev Address Sig for %s: %w", address, err)
	}
	defer resp.Body.Close()

	revSig, err := io.ReadAll(resp.Body)

	return string(revSig), err
}

func (c *Client) Suggest(name string, options *SuggestOptions) ([]models.Record, error) {
	var records []models.Record
	query, _ := url.Parse(baseURL + "nfd/suggest/" + name)

	if options != nil {
		query.RawQuery = options.ToQueryParams().Encode()
	}

	resp, err := c.get(query.String(), nil)
	if err != nil {
		return records, fmt.Errorf("Unable to get suggestions for %s: %w", name, err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&records)

	return records, err
}

func (c *Client) get(url string, header map[string]string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	for key, value := range header {
		request.Header.Add(key, value)
	}

	//TODO: Better error handling
	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("HTTP Status Code: %s", resp.Status)
	}

	return resp, nil
}

type View string

const (
	BRIEF_VIEW     View = "brief"
	TINY_VIEW      View = "tiny"
	THUMBNAIL_VIEW View = "thumbnail"
	FULL_VIEW      View = "full"
)

type NameOptions struct {
	View        View
	Poll        bool
	NoCache     bool
	IfNoneMatch string
}

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

type ActivityOptions struct {
	Type        Type
	After       time.Time
	Limit       int
	Sort        SortOption
	IfNoneMatch string
}

type Event = models.Event
type Category = models.Category
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

type AuctionOptions struct {
	Name         string
	StartingSoon bool
}
