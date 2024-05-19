package postgres

import (
	"context"
	"github.com/danyaobertan/exchangemonitor/models"
	"time"
)

const ConnectionTimeout = 5 * time.Second

func (p Postgres) AddSubscription(ctx context.Context, subscriber models.Subscriber) error {
	ctx, cancelFunc := context.WithTimeout(ctx, ConnectionTimeout)
	defer cancelFunc()

	_, err := p.connectionPool.Exec(ctx, "INSERT INTO subscribers (email) VALUES ($1)", subscriber.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p Postgres) GetSubscription(ctx context.Context, subscriber models.Subscriber) (models.Subscriber, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, ConnectionTimeout)
	defer cancelFunc()

	var sub models.Subscriber
	if err := p.connectionPool.QueryRow(ctx, "SELECT email FROM subscribers WHERE email = $1", subscriber.Email).Scan(&sub.Email); err != nil {
		return models.Subscriber{}, err
	}

	return sub, nil
}
