package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/sample-catalog/pkg/api/v1/items"
)


func RegisterRoutes(app *fiber.App){
	v1items := items.NewHandler()
	routes := app.Group("/items")
	routes.Post("/items-inquiry", v1items.GetItemsCatalog)
}