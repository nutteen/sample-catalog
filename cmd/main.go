package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nutteen/png-core/core/logger"
	middlewarelogger "github.com/nutteen/png-core/core/middlewares/logger"
	"github.com/nutteen/png-core/core/validator"
	"github.com/nutteen/png-core/core/middlewares/usercontext"
	"github.com/nutteen/sample-catalog/pkg/router"
)

func main() {
	//config.LoadConfig()

	logger.InitializeLogger(logger.LoggerConfig{
		IsProductionMode: true,
	})

	validator.InitializeValidator()

	app := fiber.New()
	//dbInstance := db.Init(config.AppConfig.DBUrl)

	// Setup middlewares
	app.Use(requestid.New())
	app.Use(usercontext.New())
	app.Use(middlewarelogger.NewLogger(logger.Log, middlewarelogger.ConfigDefault))

	// Setup application service
	//inviteRepository := infrastructure.NewGormInviteRepository(dbInstance)
	//appService := application.NewAppService(inviteRepository)

	// Registers routes
	router.RegisterRoutes(app)

	//app.Listen(config.AppConfig.Port)
	app.Listen(":3000")
}