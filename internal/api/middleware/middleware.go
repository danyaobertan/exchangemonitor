package middleware

import (
	p "github.com/danyaobertan/exchangemonitor/internal/db"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
)

type Middleware struct {
	log      logger.Logger
	dbClient p.Postgres
}

func NewMiddleware(log logger.Logger, dbClient p.Postgres) *Middleware {
	return &Middleware{
		log:      log,
		dbClient: dbClient,
	}
}
