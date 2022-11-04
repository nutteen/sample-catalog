package catalogservice

import (
	"context"
	"fmt"
	"github.com/nutteen/png-core/core/logger"
	"github.com/nutteen/sample-catalog/pkg/models"
	"github.com/nutteen/sample-catalog/pkg/repository/items"
)

type Service struct {
	itemRepository items.ItemRepository
}

func NewService(itemRepository items.ItemRepository) *Service {
	return &Service{
		itemRepository: itemRepository,
	}
}

func (service Service) GetItemsCatalog (ctx context.Context, request models.GetItemsRequest) (*models.GetItemsResponse, error){
	logger.Log.Debug(fmt.Sprintf("Getting items"))

	itemEntities, err := service.itemRepository.GetByIds(request.ItemIds)

	// Return data
	var itemsDto []models.ItemDetail

	for _, item := range itemEntities {
		itemDto := models.ItemDetail{
			ItemId: item.ID,
			Description: item.Description,
		}
		itemsDto = append(itemsDto, itemDto)
	}

	response := models.GetItemsResponse{
		Items: itemsDto,
	}

	return &response, err
}