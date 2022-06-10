package models

import "time"

type Record struct {
	AppID            int        `json:"appID"`
	ASAID            int        `json:"asaID"`
	CaAlgo           []string   `json:"caAlgo"`
	Category         string     `json:"category"`
	CurrentAsofBlock int        `json:"currentAsOfBlock"`
	MatchCheck       string     `json:"match-check"`
	MetaTags         []string   `json:"metaTags"`
	Name             string     `json:"name"`
	NFDAccount       string     `json:"nfdAccount"`
	Owner            string     `json:"owner"`
	Properties       Properties `json:"properties"`
	ReservedFor      string     `json:"reservedFor"`
	SaleType         SaleType   `json:"saleType"`
	SellAmount       int        `json:"sellAmount"`
	Seller           string     `json:"seller"`
	SigNameAddress   string     `json:"sigNameAddress"`
	State            State      `json:"state"`
	Tags             []string   `json:"tags"`
	TimeChanged      time.Time  `json:"timeChanged"`
	TimeCreated      time.Time  `json:"timeCreated"`
	TimePurchased    time.Time  `json:"timePurchased"`
	UnverifiedCaAlgo []string   `json:"unverifiedCaAlgo"`
}
