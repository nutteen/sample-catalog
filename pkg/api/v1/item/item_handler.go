package item

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/png-core/core/utils/errorutils"
	corevalidator "github.com/nutteen/png-core/core/validator"
	"github.com/nutteen/sample-catalog/pkg/model"
	"github.com/nutteen/sample-catalog/pkg/service/catalogservice"
)

type Handler struct {
	catalogService 	*catalogservice.Service
}

func NewHandler(service *catalogservice.Service) *Handler {
	return &Handler{catalogService: service}
}

func (h Handler) GetItemsCatalog(c *fiber.Ctx) error {
	request := model.GetItemsRequest{}

	ctx	:= c.UserContext()

	if err := c.BodyParser(&request); err != nil {
		errorResponse := errorutils.BAD_REQUEST_BODY_PARSE_ERROR.NewErrorResponseModel(err.Error())
		return c.Status(errorResponse.Error.Status).JSON(errorResponse)
	}

	if err := corevalidator.Validate.Struct(&request); err != nil {
		invalidParams := errorutils.ValidationErrorsToInvalidParams(err.(validator.ValidationErrors))
		errorResponse := errorutils.BAD_REQUEST_VALDATION_ERROR.NewErrorResponseModel(err.Error(), invalidParams)
		return c.Status(errorResponse.Error.Status).JSON(errorResponse)
	}

	response, err := h.catalogService.GetItemsCatalog(ctx, request)

	if err != nil {
		errorResponse := errorutils.UNKNOWN_ERROR.NewErrorResponseModel(err.Error())
		return c.Status(errorResponse.Error.Status).JSON(errorResponse)
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}