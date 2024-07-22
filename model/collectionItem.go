package model

type CollectionItem struct {
	ID            string      `json:"id"`
	CmsLocaleId   string      `json:"cmsLocaleId"`
	LastPublished string      `json:"lastPublished"`
	LastUpdated   string      `json:"lastUpdated"`
	CreatedOn     string      `json:"createdOn"`
	IsArchived    bool        `json:"isArchived"`
	IsDraft       bool        `json:"isDraft"`
	FieldData     interface{} `json:"fieldData"`
}

type CollectionItemsResponse struct {
	Items      []CollectionItem `json:"items"`
	Pagination Pagination       `json:"pagination"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
