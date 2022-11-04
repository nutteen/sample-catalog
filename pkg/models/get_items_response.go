package models

type GetItemsResponse struct {
	Items 		[]ItemDetail 		`json:"items"`
}

type ItemDetail struct {
	ItemId	string	`json:"itemId"`
	Description string `json:"description"`
}