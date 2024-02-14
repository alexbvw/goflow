package model

type Collection struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	SingularName string `json:"singularName"`
	Slug         string `json:"slug"`
	CreatedOn    string `json:"createdOn"`
	LastUpdated  string `json:"lastUpdated"`
}

type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}
