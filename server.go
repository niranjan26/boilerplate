package postman

import (
	"fmt"
	"net/http"
	"time"

	"postman/internal/api/v1/handlers"
	"postman/internal/api/v1/middlewares"
	"postman/internal/config"
	"postman/internal/service"
	"postman/internal/storage"

	"github.com/gorilla/mux"
)

type Handler interface {
	AddRoutes(router *mux.Router)
}

func registerHandlers(router *mux.Router, mhandlers ...Handler) {
	for _, handler := range mhandlers {
		handler.AddRoutes(router)
	}
}

func NewServer(config config.Config, dbStorage *storage.DBStorage) *http.Server {
	router := mux.NewRouter()

	router.Use(middlewares.PopulateContextMiddleware())

	postmanService := service.NewPostmanService(dbStorage)
	postmanHandler := handlers.NewPostmanHandler(postmanService)

	registerHandlers(router,
		&handlers.HealthCheck{},
		postmanHandler,
	)

	address := fmt.Sprintf("%s:%s", config.Host, config.Port)

	return &http.Server{
		Handler: router,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
