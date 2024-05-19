package main

import (
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/api"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/internal/db/postgres"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Load the configuration
	conf := config.InitConfig()
	fmt.Printf("Configuration for %v loaded successfully\n", conf.App.Name)
	// Create a logger
	log := logger.InitLogger(conf.App.Debug)
	log.Info("Logger created successfully")

	// Create a connection pool to the database
	connectionPool := postgres.InitDB(log, conf.DB)
	defer connectionPool.Close()
	postgresClient := postgres.NewPostgres(connectionPool, log)

	// Apply migrations
	postgres.RunMigrations(log, conf.DB)

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	shutDownChannel := make(chan struct{})
	waitGroup := &sync.WaitGroup{}
	// Start the application
	go api.Run(*postgresClient, conf, log, shutDownChannel, waitGroup)

	<-quitChannel
	close(shutDownChannel)
	waitGroup.Wait()
	log.Info("Server shutdown")
}
