package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"time"
)

type Postgres struct {
	connectionPool *pgxpool.Pool
	log            logger.Logger
}

func NewPostgres(connectionPool *pgxpool.Pool, log logger.Logger) *Postgres {
	return &Postgres{
		connectionPool: connectionPool,
		log:            log,
	}
}

func Config(log logger.Logger, conf config.DBConfiguration) *pgxpool.Config {
	dbConfig, err := pgxpool.ParseConfig(conf.ConnectionURL)
	if err != nil {
		log.Fatal("Error parsing connection URL" + err.Error())
	}

	dbConfig.MaxConns = int32(conf.DefaultMaxConnections)
	dbConfig.MaxConnLifetime = time.Duration(conf.DefaultMaxConnectionLifetime) * time.Second
	dbConfig.MaxConnIdleTime = time.Duration(conf.DefaultMaxConnectionIdleTime) * time.Second
	dbConfig.HealthCheckPeriod = time.Duration(conf.DefaultHealthCheckPeriod) * time.Second
	dbConfig.ConnConfig.ConnectTimeout = time.Duration(conf.DefaultHealthCheckTimeout) * time.Second

	dbConfig.BeforeAcquire = func(_ context.Context, _ *pgx.Conn) bool {
		log.Info("Acquiring connection")
		return true
	}

	dbConfig.AfterRelease = func(_ *pgx.Conn) bool {
		log.Info("Releasing connection")
		return true
	}

	dbConfig.BeforeClose = func(_ *pgx.Conn) {
		log.Info("Closing connection")
	}

	return dbConfig
}

func RunMigrations(log logger.Logger, conf config.DBConfiguration) {
	fmt.Println("Migration source:", conf.MigrationSource)
	fmt.Println("Full migration path:", conf.ConnectionURL+conf.MigrationQueryParams)
	m, err := migrate.New(
		conf.MigrationSource,
		conf.ConnectionURL+conf.MigrationQueryParams,
	)
	if err != nil {
		log.Fatal("Error creating migration instance" + err.Error())
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Error running migrations" + err.Error())
	}

	log.Info("Migrations ran successfully")
}

func InitDB(log logger.Logger, conf config.DBConfiguration) *pgxpool.Pool {
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config(log, conf))
	if err != nil {
		log.Fatal("Error creating connection pool" + err.Error())
	}
	err = connPool.Ping(context.Background())
	if err != nil {
		log.Fatal("Error pinging database" + err.Error())
	}
	log.Info("Database connection established")
	return connPool
}
