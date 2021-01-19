package bullet

type Device struct {
	ID           string `json:"iden"`
	Active       bool   `json:"active"`
	Icon         string `json:"icon"`
	Nickname     string `json:"nickname"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	AppVersion   int    `json:"app_version"`
}

type Devices struct {
	Items []Device `json:"devices"`
}
