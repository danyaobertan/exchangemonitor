package handler

import (
	"github.com/danyaobertan/exchangemonitor/internal/api/middleware"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/internal/db"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	conf     *config.Configuration
	log      logger.Logger
	dbClient *db.Postgres
	middleware.Middleware
	//	subsvriber *subscriber.Subscriber
}

func NewHandler(conf *config.Configuration, log logger.Logger, dbClient *db.Postgres) *Handler {
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
