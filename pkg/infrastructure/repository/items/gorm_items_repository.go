package items

import (
	"github.com/nutteen/sample-catalog/pkg/domain/items"
	"gorm.io/gorm"
)

type GormItemRepository struct {
	db *gorm.DB
}

func NewGormItemRepository(gormDB *gorm.DB) *GormItemRepository{
	return &GormItemRepository{db: gormDB}
}

func (repo GormItemRepository) GetByIds(ids []string) ([]items.Item, error) {
	var itemEntities []items.Item
	result := repo.db.Find(itemEntities, ids)
	return itemEntities, result.Error
}