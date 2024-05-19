package api

import (
	"context"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	p "github.com/danyaobertan/exchangemonitor/internal/db/postgres"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/danyaobertan/exchangemonitor/internal/subscriber"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	WriteTimeout = 60
	ReadTimeout  = 60
	IdleTimeout  = 60

	GracefulShutdownTimeout = 5
)

func Run(dbClient p.Postgres, conf *config.Configuration, log logger.Logger, shutDownChannel chan struct{}, wg *sync.WaitGroup) {
	wg.Add(2) // Adding two because we now have two goroutines: HTTP server and email worker

	// Start the email worker
	go subscriber.StartEmailWorker(dbClient, conf, log, wg, shutDownChannel)

	// Configure and start the HTTP server
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(conf.App.Port),
		WriteTimeout: WriteTimeout * time.Second,
		ReadTimeout:  ReadTimeout * time.Second,
		IdleTimeout:  IdleTimeout,
		Handler:      getRouter(dbClient, conf, log),
	}

	httpErr := make(chan error, 1)
	go func() {
		httpErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-httpErr:
		log.Error("Error starting server: %s" + err.Error())
		return
	case <-shutDownChannel:
		log.Info("Shutting down server")
		break
	}

	ctx, cancel := context.WithTimeout(context.Background(), GracefulShutdownTimeout*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Error("Error shutting down server: %s" + err.Error())
	}
	log.Info("Graceful shutdown complete")
}
