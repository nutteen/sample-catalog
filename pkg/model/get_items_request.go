package model

type GetItemsRequest struct {
	ItemIds 		[]string 		`json:"itemIds" validate:"required"`
}