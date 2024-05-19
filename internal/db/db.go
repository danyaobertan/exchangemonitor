package db

import (
	"context"
	"github.com/danyaobertan/exchangemonitor/models"
)

type Database interface {
	AddSubscription(ctx context.Context, subscriber models.Subscriber) error
	GetSubscription(ctx context.Context, subscriber models.Subscriber) (models.Subscriber, error)
	GetAllSubscriptions(ctx context.Context) ([]models.Subscriber, error)
}
