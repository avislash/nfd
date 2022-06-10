package models

import "time"

type AnalyticRecord struct {
	Data      AnalyticEvent `json:"data"`
	Price     int           `json:"price"`
	Timestamp time.Time     `json:"timeStamp"`
}

type AnalyticRecords struct {
	//MatchCheck string `json:"match-check"` Not Returned
	Results []AnalyticRecord `json:"results"`
	Total   int              `json:"total"`
}
