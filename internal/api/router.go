package api

import (
	h "github.com/danyaobertan/exchangemonitor/internal/api/handler"
	m "github.com/danyaobertan/exchangemonitor/internal/api/middleware"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	p "github.com/danyaobertan/exchangemonitor/internal/db"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/go-chi/chi/v5"
)

func getRouter(dbClient p.Postgres, conf *config.Configuration, log logger.Logger) *chi.Mux {
	router := chi.NewRouter()

	handler := h.NewHandler(conf, log, &dbClient)
	middleware := m.NewMiddleware(log, dbClient)

	router.Use(middleware.Logger)
	handler.SetupRoutes(router)
	return router
}
