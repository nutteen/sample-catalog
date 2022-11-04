package items

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/png-core/core/logger"
	"github.com/nutteen/png-core/core/utils/errorutils"
	"github.com/nutteen/png-core/core/validator"
	"github.com/nutteen/sample-catalog/pkg/models"
	"github.com/nutteen/sample-catalog/pkg/service/catalogservice"
)

type Handler struct {
	catalogService 	*catalogservice.Service
}

func NewHandler(service *catalogservice.Service) *Handler {
	return &Handler{catalogService: service}
}

func (h Handler) GetItemsCatalog(c *fiber.Ctx) error {
	request := models.GetItemsRequest{}

	ctx	:= c.UserContext()
	logger.Log.Debug("Get item Catalog handler", logger.LogContext(ctx))

	if err := c.BodyParser(&request); err != nil {
		errorResponse := errorutils.NewErrorResponseModel(errorutils.BAD_REQUEST_BODY_PARSE_ERROR, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	if err := validator.Validate.Struct(&request); err != nil {
		errorResponse := errorutils.NewErrorResponseModel(errorutils.BAD_REQUEST_VALDATION_ERROR, err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse)
	}

	response, err := h.catalogService.GetItemsCatalog(ctx, request)

	if err != nil {
		errorResponse := errorutils.NewErrorResponseModel(errorutils.GENERIC_ERROR, err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse)
	}

	return c.Status(fiber.StatusCreated).JSON(&response)
}