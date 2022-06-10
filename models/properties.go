package models

type Properties struct {
	Internal    map[string]string `json:"internal"`
	UserDefined map[string]string `json:"userDefined"`
	Verified    map[string]string `json:"verified"`
}
