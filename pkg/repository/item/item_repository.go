package item

import "github.com/nutteen/sample-catalog/pkg/domain/item"

type ItemRepository interface {
	GetByIds(ids []string) ([]item.Item, error)
}