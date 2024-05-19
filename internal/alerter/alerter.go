package alerter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/danyaobertan/exchangemonitor/internal/config"
	p "github.com/danyaobertan/exchangemonitor/internal/db/postgres"
	"github.com/danyaobertan/exchangemonitor/internal/logger"
	"github.com/danyaobertan/exchangemonitor/models"
	"github.com/danyaobertan/exchangemonitor/pkg/currency"
	"github.com/danyaobertan/exchangemonitor/pkg/email"
	"go.uber.org/zap"
)

const AlertPeriod = 24 * time.Hour

func StartEmailWorker(dbClient p.Postgres, config *config.Configuration, log logger.Logger, wg *sync.WaitGroup, shutDownChannel chan struct{}) {
	defer wg.Done()

	sendEmails := func() {
		subs, err := GetSubscribers(dbClient)

		if err != nil {
			log.Error("Failed to get subscribers: ", zap.Error(err))
			return
		}

		rate, err := GetExchangeRate()
		if err != nil {
			log.Error("Failed to get exchange rate: ", zap.Error(err))
			return
		}

		for _, sub := range subs {
			emailData := models.EmailDataObject{
				Name:    sub.Email,
				Subject: "Daily Currency Exchange Update",
				Message: fmt.Sprintf("Current USD to UAH exchange rate is %f", rate),
			}

			err = email.SendEmail(config.SMTP, sub.Email, emailData)
			if err != nil {
				log.Error("Failed to send email: ", zap.Error(err))
			} else {
				log.Info(fmt.Sprintf("Email sent successfully to %s", sub.Email))
			}
		}
	}

	// Send immediately upon startup
	sendEmails()

	// Set up a ticker to send emails daily
	ticker := time.NewTicker(AlertPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sendEmails()
		case <-shutDownChannel:
			log.Info("Shutting down email worker")
			return
		}
	}
}

func GetSubscribers(dbClient p.Postgres) ([]models.Subscriber, error) {
	ctx := context.Background()

	subs, err := dbClient.GetAllSubscriptions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %s", err)
	}

	return subs, nil
}

func GetExchangeRate() (float64, error) {
	rate, err := currency.FetchCurrentRateNBU()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch rate: %s", err)
	}

	return rate, nil
}
