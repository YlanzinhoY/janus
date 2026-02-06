package module

import (
	"ylanzinhoy-operator-management/config"
	"ylanzinhoy-operator-management/internal/application/useCase"
	"ylanzinhoy-operator-management/internal/infrastructure/persistence/storage"
	"ylanzinhoy-operator-management/internal/infrastructure/web"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Api struct {
	app *fiber.App
	cfg *config.Config
}

func NewApi() *Api {
	return &Api{
		app: fiber.New(),
		cfg: config.GetInstance(),
	}

}

func (a *Api) ApiModuleInitialize(database *mongo.Collection) {

	webHandler := a.Handler(database)
	a.Routes(webHandler)

}

func (a *Api) Handler(database *mongo.Collection) *web.WebHandlerAdapter {
	storageOperator := storage.NewAdapter(database.Name(), "operator", database)
	webHandler := web.NewWebAdapter(useCase.NewOperatorUseCase(storageOperator))

	return webHandler
}

func (a *Api) Routes(handler *web.WebHandlerAdapter) {

	a.app.Post("/operator", handler.PostOperator)
	a.app.Get("/operator", handler.GetOperators)
	a.app.Get("/operator/:id", handler.GetOperatorById)
	a.app.Put("/operator/:id", handler.UpdateOperator)
	a.app.Delete("/operator/:id", handler.DeleteOperator)
	a.app.Put("/operator/change-password", handler.ChangePassword)

}

func (a *Api) Start() error {
	err := a.app.Listen(":3000")
	if err != nil {
		return err
	}

	return nil
}
