package model

type CollectionItem struct {
	ID            string    `json:"id"`
	CmsLocaleId   string    `json:"cmsLocaleId"`
	LastPublished string    `json:"lastPublished"`
	LastUpdated   string    `json:"lastUpdated"`
	CreatedOn     string    `json:"createdOn"`
	IsArchived    bool      `json:"isArchived"`
	IsDraft       bool      `json:"isDraft"`
	FieldData     FieldData `json:"fieldData"`
}

type FieldData struct {
	Featured       bool   `json:"featured"`
	Color          string `json:"color"`
	Name           string `json:"name"`
	PostBody       string `json:"post-body"`
	PostSummary    string `json:"post-summary"`
	MainImage      Image  `json:"main-image"`
	ThumbnailImage Image  `json:"thumbnail-image"`
	Slug           string `json:"slug"`
}

type Image struct {
	FileId string  `json:"fileId"`
	URL    string  `json:"url"`
	Alt    *string `json:"alt"` // Use a pointer to handle null values
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
