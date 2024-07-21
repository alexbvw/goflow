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
	DateAndTime      string `json:"date-and-time"`
	TicketsLink      string `json:"tickets-link"`
	Name             string `json:"name"`
	City             string `json:"city"`
	Country          string `json:"country"`
	Location         string `json:"location"`
	Venue            string `json:"venue"`
	Description      string `json:"description"`
	ShortDescription string `json:"short-description"`
	Image            Image  `json:"image"`
	Slug             string `json:"slug"`
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
