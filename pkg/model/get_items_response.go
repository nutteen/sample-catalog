package model

import "github.com/shopspring/decimal"

type GetItemsResponse struct {
	Items 		[]ItemDetail 		`json:"items"`
}

type ItemDetail struct {
	ItemId	string	`json:"itemId"`
	Description string `json:"description"`
	Price decimal.Decimal `json:"price"`
}