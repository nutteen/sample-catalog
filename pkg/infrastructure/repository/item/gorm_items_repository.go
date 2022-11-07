package item

import (
	"github.com/nutteen/sample-catalog/pkg/domain/item"
	"gorm.io/gorm"
)

type GormItemRepository struct {
	db *gorm.DB
}

func NewGormItemRepository(gormDB *gorm.DB) *GormItemRepository{
	return &GormItemRepository{db: gormDB}
}

func (repo GormItemRepository) GetByIds(ids []string) ([]item.Item, error) {
	var itemEntities []item.Item
	result := repo.db.Find(&itemEntities, ids)
	return itemEntities, result.Error
}