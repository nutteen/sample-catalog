package items

import "github.com/nutteen/sample-catalog/pkg/domain/items"

type ItemRepository interface {
	GetByIds(ids []string) ([]items.Item, error)
}