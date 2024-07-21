package model

type Site struct {
	ID            string   `json:"id"`
	WorkspaceId   string   `json:"workspaceId"`
	DisplayName   string   `json:"displayName"`
	ShortName     string   `json:"shortName"`
	PreviewUrl    string   `json:"previewUrl"`
	TimeZone      string   `json:"timeZone"`
	CreatedOn     string   `json:"createdOn"`
	LastUpdated   string   `json:"lastUpdated"`
	LastPublished string   `json:"lastPublished"`
	CustomDomains []string `json:"customDomains"`
	Locales       Locales  `json:"locales"`
}

type Locales struct {
	Primary   Locale   `json:"primary"`
	Secondary []string `json:"secondary"`
}

type Locale struct {
	ID           string `json:"id"`
	CmsLocaleId  string `json:"cmsLocaleId"`
	Enabled      bool   `json:"enabled"`
	DisplayName  string `json:"displayName"`
	Redirect     bool   `json:"redirect"`
	Subdirectory string `json:"subdirectory"`
	Tag          string `json:"tag"`
}

type SitesResponse struct {
	Sites []Site `json:"sites"`
}

type SiteResponse struct {
	Site Site `json:"site"`
}
