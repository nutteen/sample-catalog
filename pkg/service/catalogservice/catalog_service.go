package catalogservice

import (
	"context"
	"fmt"
	"github.com/nutteen/png-core/core/logger"
	"github.com/nutteen/sample-catalog/pkg/model"
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

func (service Service) GetItemsCatalog (ctx context.Context, request model.GetItemsRequest) (*model.GetItemsResponse, error){
	logger.Log.Debug(fmt.Sprintf("Getting items"))

	itemEntities, err := service.itemRepository.GetByIds(request.ItemIds)

	// Return data
	var itemsDto []model.ItemDetail

	for _, item := range itemEntities {
		itemDto := model.ItemDetail{
			ItemId: item.ID,
			Description: item.Description,
			Price: item.Price,
		}
		itemsDto = append(itemsDto, itemDto)
	}

	response := model.GetItemsResponse{
		Items: itemsDto,
	}

	return &response, err
}