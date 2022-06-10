package models

type Properties struct {
	Internal    map[string]string `json:"internal"`
	UserDefined map[string]string `json:"userDefined"`
	Verified    map[string]string `json:"verified"`
}

/*func (p *Properties) UnmarshalJSON(data []byte) error {

	s := struct {
		Internal    map[string]string `json:"internal"`
		UserDefined map[string]string `json:"userDefined"`
		Verified    map[string]string `json:"verified"`
	}{
		Internal:    make(map[string]string),
		UserDefined: make(map[string]string),
		Verified:    make(map[string]string),
	}
	err := json.Unmarshal(data, &s)

	p.Internal = s.Internal
	p.UserDefined = s.UserDefined
	p.Verified = s.Verified

	return err
}*/
