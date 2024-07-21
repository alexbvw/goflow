package model

// Struct to match the overall response structure
type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}

type CollectionResponse struct {
	Collection Collection `json:"collection"`
}

// Struct to match the Collection object
type Collection struct {
	ID           string        `json:"id"`
	DisplayName  string        `json:"displayName"`
	SingularName string        `json:"singularName"`
	Slug         string        `json:"slug"`
	CreatedOn    string        `json:"createdOn"`
	LastUpdated  string        `json:"lastUpdated"`
	Fields       []interface{} `json:"fields"`
}
