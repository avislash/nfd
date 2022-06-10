package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/avislash/nfd/models"
)

const (
	MAIN_DOMAIN    string = "https://api.nf.domains/"
	TESTNET_DOMAIN string = "https://api.testnet.nf.domains/"
	BETA_DOMAIN    string = "https://api.betanet.nf.domains/"
)

type Client struct {
	*http.Client
	root string
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{
		&http.Client{},
		MAIN_DOMAIN,
	}

	for _, applyOption := range options {
		applyOption(client)
	}

	return client
}

func (c *Client) Version() (models.Version, error) {
	var version models.Version
	resp, err := c.get("info/version", nil)
	if err != nil {
		return version, fmt.Errorf("Error fetching version: %w", err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&version)
	return version, err
}

//TODO: Find out what if-none-match is
func (c *Client) Name(name string, options *NameOptions) (models.Record, error) {
	endPoint := "nfd/" + name
	var record models.Record

	resp, err := c.get(endPoint, options)
	if err != nil {
		return record, fmt.Errorf("Unable to fetch NFD Record for %s: %w", name, err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&record)

	return record, err
}

func (c *Client) Activity(names []string, options *ActivityOptions) ([]models.ActivityRecord, error) {
	var activityRecords []models.ActivityRecord

	_names := make(url.Values)
	for _, name := range names {
		_names.Add("name", name)
	}
	endPoint := "nfd/activity?" + _names.Encode()

	resp, err := c.get(endPoint, options)
	if err != nil {
		return activityRecords, fmt.Errorf("Unable to get Activity Records for %s: %w", names, err)
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&activityRecords)

	return activityRecords, err

}

func (c *Client) Address(addresses []string, options *AddressOptions) ([]models.Record, error) {
	var records []models.Record

	_addresses := make(url.Values)
	for _, address := range addresses {
		_addresses.Add("address", address)
	}
	endPoint := "nfd/address?" + _addresses.Encode()

	resp, err := c.get(endPoint, options)
	if err != nil {
		return records, fmt.Errorf("Unable to get Records for %s: %w", addresses, err)
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&records)

	return records, err
}

func (c *Client) Analytics(options *AnalyticsOptions) (models.AnalyticRecords, error) {
	var records models.AnalyticRecords
	endPoint := "nfd/analytics"

	resp, err := c.get(endPoint, options)
	if err != nil {
		return records, fmt.Errorf("Unable to get Analytics Records: %w", err)
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&records)

	return records, err
}

func (c *Client) Auction(options *AuctionOptions) ([]models.AuctionAndPrice, error) {
	var openAuctions []models.AuctionAndPrice
	endPoint := "nfd/auction"

	resp, err := c.get(endPoint, options)
	if err != nil {
		return openAuctions, fmt.Errorf("Unable to get Auction and Price: %w", err)
	}
	defer resp.Close()
	err = json.NewDecoder(resp).Decode(&openAuctions)

	return openAuctions, err
}

func (c *Client) Browse(options *BrowseOptions) ([]models.Record, error) {
	var records []models.Record
	endPoint := "nfd/browse"

	resp, err := c.get(endPoint, options)
	if err != nil {
		return records, fmt.Errorf("Unable to Browse NFD: %w", err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&records)

	return records, nil

}

func (c *Client) IsValid(appID int) (models.Valid, error) {
	var isValid models.Valid
	endPoint := "nfd/isValid/" + strconv.Itoa(appID)

	resp, err := c.get(endPoint, nil)
	if err != nil {
		return isValid, fmt.Errorf("Unable check Validity for %d: %w", appID, err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&isValid)

	return isValid, err
}

func (c *Client) NameSig(name string) (string, error) {
	var nameSig string
	endPoint := "nfd/nameSig/" + name

	resp, err := c.get(endPoint, nil)
	if err != nil {
		return "", fmt.Errorf("Unable to Get Name Sig for %s: %w", name, err)
	}
	defer resp.Close()

	rawNameSig, err := io.ReadAll(resp)

	if err == nil {
		//Returned value contains a formatted string which contains quote and a new line.
		//Trim the new line and unquote the string so it can be reused directly as an address
		//without requring the user to reformat it.
		nameSig, err = strconv.Unquote(strings.TrimSuffix(string(rawNameSig), "\n"))
	}
	return nameSig, err
}

func (c *Client) RevSig(address string) (string, error) {
	var revSig string
	endPoint := "nfd/revAddressSig/" + address

	resp, err := c.get(endPoint, nil)
	if err != nil {
		return "", fmt.Errorf("Unable to Get Rev Address Sig for %s: %w", address, err)
	}
	defer resp.Close()

	rawRevSig, err := io.ReadAll(resp)

	if err == nil {
		//Returned value contains a formatted string which contains quote and a new line.
		//Trim the new line and unquote the string so it can be reused directly as an address
		//without requring the user to reformat it.
		revSig, err = strconv.Unquote(strings.TrimSuffix(string(rawRevSig), "\n"))
	}

	return revSig, err
}

func (c *Client) Suggest(name string, options *SuggestOptions) ([]models.Record, error) {
	var records []models.Record
	endPoint := "nfd/suggest/" + name

	resp, err := c.get(endPoint, options)
	if err != nil {
		return records, fmt.Errorf("Unable to get suggestions for %s: %w", name, err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&records)

	return records, err
}

func (c *Client) Totals() (models.Totals, error) {
	var totals models.Totals
	endPoint := "nfd/totals"

	resp, err := c.get(endPoint, nil)
	if err != nil {
		return totals, fmt.Errorf("Error fetching totals: %w", err)
	}
	defer resp.Close()

	err = json.NewDecoder(resp).Decode(&totals)

	return totals, err
}

func (c *Client) get(endPoint string, options option) (io.ReadCloser, error) {
	var header map[string]string

	fmt.Println("endPoint in Get: ", endPoint)
	query, err := url.Parse(c.root + endPoint)
	if err != nil {
		return nil, err
	}

	fmt.Println("Query Pre OptionsParsing: ", query.String())

	if nil != options {
		params := options.GetQueryParams().Encode()
		if len(query.RawQuery) != 0 { //Endpoint may have come with required parameters
			query.RawQuery = query.RawQuery + "&" + params
		} else {
			query.RawQuery = params
		}

		header = options.GetHeader()
	}

	fmt.Println("Query: ", query.String())

	request, err := http.NewRequest("GET", query.String(), nil)
	for key, value := range header {
		request.Header.Add(key, value)
	}

	resp, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	//TODO: Make this error message a bit better.
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("HTTP Status Code: %s", resp.Status)
	}

	return resp.Body, nil
}
