package handler

import (
	"github.com/danyaobertan/exchangemonitor/internal/api/middleware"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/internal/db/postgres"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	conf     *config.Configuration
	log      logger.Logger
	dbClient *postgres.Postgres
	middleware.Middleware
}

func NewHandler(conf *config.Configuration, log logger.Logger, dbClient *postgres.Postgres) *Handler {
	return &Handler{
		conf:     conf,
		log:      log,
		dbClient: dbClient,
	}
}

func (h *Handler) SetupRoutes(router *chi.Mux) {
	router.Get("/rate", h.HandleGetRate)
	router.Post("/subscribe", h.HandleSubscribe)
}
