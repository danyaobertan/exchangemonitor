package alerter

import (
	"context"
	"errors"
	"fmt"
	"github.com/danyaobertan/exchangemonitor/internal/config"
	p "github.com/danyaobertan/exchangemonitor/internal/db/postgres"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/danyaobertan/exchangemonitor/models"
	"github.com/danyaobertan/exchangemonitor/pkg/currency"
	"github.com/danyaobertan/exchangemonitor/pkg/email"
	"go.uber.org/zap"
	"sync"
	"time"
)

const AlertPeriod = 24 * time.Hour

func StartEmailWorker(dbClient p.Postgres, config *config.Configuration, log logger.Logger, wg *sync.WaitGroup, shutDownChannel chan struct{}) {
	defer wg.Done()
	ticker := time.NewTicker(AlertPeriod)

	for {
		select {
		case <-ticker.C:
			subs, err := GetSubscribers(dbClient)
			if err != nil {
				log.Error("Failed to get subscribers: ", zap.Error(err))
				continue
			}
			rate, err := GetExchangeRate()
			if err != nil {
				log.Error("Failed to get exchange rate: ", zap.Error(err))
				continue
			}
			for _, sub := range subs {
				err = email.SendEmail(config.SMTP, sub.Email, "Daily Update", fmt.Sprintf("Here's your daily update. Current USD to UAH exchange rate is %f", rate))
				if err != nil {
					log.Error("Failed to send email: ", zap.Error(err))
				} else {
					log.Info(fmt.Sprintf("Email sent successfully to %s", sub.Email))
				}
			}
		case <-shutDownChannel:
			log.Info("Shutting down email worker")
			ticker.Stop()
			return
		}
	}
}

func GetSubscribers(dbClient p.Postgres) ([]models.Subscriber, error) {
	ctx := context.Background()
	subs, err := dbClient.GetAllSubscriptions(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get subscribers: %s", err))
	}
	return subs, nil
}

func GetExchangeRate() (float64, error) {
	rate, err := currency.FetchCurrentRateNBU()
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Failed to fetch rate: %s", err))
	}
	return rate, nil
}