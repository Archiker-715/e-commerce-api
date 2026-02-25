package app

import (
	"log"
	"net/http"

	"github.com/Archiker-715/e-commerce-api/internal/repo/pg"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

type app struct {
	DB     *gorm.DB
	Router *mux.Router
}

func Run() {
	// TODO: переписать запуск приложения  учетом зависимостей
	newApp().InitServices()
}

func newApp() *app {
	db := pg.Connect()

	r := mux.NewRouter()
	r.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s %s", r.Method, r.URL.Path)
			h.ServeHTTP(w, r)
		})
	})
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router := r.PathPrefix("/api/v1").Subrouter()

	app := app{
		DB:     db,
		Router: router,
	}

	return &app
}
