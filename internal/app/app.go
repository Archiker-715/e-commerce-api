package app

import (
	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type app struct {
	DB     *gorm.DB
	Router *mux.Router

	Repos    *repositories
	Services *services
	Handlers *handlers
}

func Run() *app {
	db := pg.Connect()
	repos := newRepositories(db)
	services := newServices(repos)
	router := newRouter()

	app := &app{
		DB:       db,
		Router:   router,
		Repos:    repos,
		Services: services,
	}

	app.initHandlers()
	app.initRoutes()
	app.startKafkaConsumers()

	return app
}
