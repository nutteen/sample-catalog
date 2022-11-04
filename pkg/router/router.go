package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/sample-catalog/pkg/api/v1/items"
	"github.com/nutteen/sample-catalog/pkg/service/catalogservice"
)


func RegisterRoutes(app *fiber.App, service *catalogservice.Service) {
	v1items := items.NewHandler(service)
	routes := app.Group("/items")
	routes.Post("/items-inquiry", v1items.GetItemsCatalog)
}