package models

type Valid struct {
	Valid          bool   `json:"isValid"`
	Message        string `json:"message"`
	Name           string `json:"name"`
	SigNameAddress string `json:"sigNameAddress"`
}
