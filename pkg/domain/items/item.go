package items

import (
	"github.com/shopspring/decimal"
)

type Item struct {
	ID				string `gorm:"primaryKey"`
	Description		string
	Price			decimal.Decimal
}
