package models

import "time"

type Auction struct {
	CeilingPrice int       `json:"ceilingPrice"`
	CurrentPrice int       `json:"int"`
	EndTime      time.Time `json:"endTime"`
	FloorPrice   int       `json:"floorPrice"`
	Name         string    `json:"name"`
	StartTime    time.Time `json:"startTime"`
}

type AuctionAndPrice struct {
	AuctionInfo     Auction `json:"auctionInfo"`
	ChangePerMinute int     `json:"changePerMinute"`
	ElapsedMinutes  int     `json:"elapsedMinutes"`
	Price           int     `json:"price"`
	TotalMinutes    int     `json:"totalMinutes"`
}
