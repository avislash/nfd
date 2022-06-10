package models

import "time"

type ActivityRecord struct {
	Block        int               `json:"block"`
	CacheControl string            `json:"cache-controle"`
	Changes      map[string]string `json:"changes"`
	ETag         string            `json:"etag"`
	//MatchCheck   string            `json:"match-check"` Not returned
	Name        string    `json:"name"`
	TimeChanged time.Time `json:"timeChanged"`
}
