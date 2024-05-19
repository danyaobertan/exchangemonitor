package main

import (
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/api"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/internal/db"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	conf := config.InitConfig()
	fmt.Printf("Configuration for %v loaded successfully\n", conf.App.Name)

	log := logger.InitLogger(conf.App.Debug)
	log.Info("Logger created successfully")

	// Create a connection pool to the database
	connectionPool := db.InitDB(log, conf.DB)
	defer connectionPool.Close()
	postgresClient := db.NewPostgres(connectionPool, log)

	// Apply migrations
	db.RunMigrations(log, conf.DB)

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	shutDownChannel := make(chan struct{})
	waitGroup := &sync.WaitGroup{}

	go api.Run(*postgresClient, conf, log, shutDownChannel, waitGroup)

	<-quitChannel
	close(shutDownChannel)
	waitGroup.Wait()
	log.Info("Server shutdown")
}