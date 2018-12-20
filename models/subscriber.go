package models

import (
	"context"
	"time"

	"github.com/go-pg/pg"

	"github.com/caosbad/ever-post-mixin-bot/session"
)

const subscribers_DDL = `
CREATE TABLE IF NOT EXISTS subscribers (
  user_id	        	VARCHAR(36) NOT NULL,
  subscriber_id     VARCHAR(36) NOT NULL,
  created_at        TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (user_id, subscriber_id)
);
`

type Subscriber struct {
	UserId        string   `sql:"user_id,pk"`
	SubscriberId string    `sql:"subscriber_id,pk"`
	CreatedAt    time.Time `sql:"created_at"`
}

func CreateSubscriber(ctx context.Context, userId, subscriberId string) (*Subscriber, error) {
	subscriber, err := FindSubscriber(ctx, userId, subscriberId)
	if err != pg.ErrNoRows {
		return subscriber, err
	}
	s := &Subscriber{
		UserId:        userId,
		SubscriberId: subscriberId,
		CreatedAt:    time.Now(),
	}
	if err := session.Database(ctx).Insert(s); err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return s, nil
}

func RemoveSubscriber(ctx context.Context, userId, subscriberId string) error {
	_, err := FindSubscriber(ctx, userId, subscriberId)
	if err != nil {
		return err
	}
	s := &Subscriber{
		UserId:        userId,
		SubscriberId: subscriberId,
	}
	if err := session.Database(ctx).Delete(s); err != nil {
		return session.TransactionError(ctx, err)
	}
	return nil
}

func FindSubscriber(ctx context.Context, userId, subscriberId string) (*Subscriber, error) {
	s := &Subscriber{
		UserId:        userId,
		SubscriberId: subscriberId,
	}
	if err := session.Database(ctx).Select(s); err != nil {
		return nil, err
	}

	return s, nil
}

func ListSubscribers(ctx context.Context, userId string) ([]*Subscriber, error) {
	var subscribers []*Subscriber
	err := session.Database(ctx).Model(&subscribers).Where("user_id = ?", userId).Select()
	if err != nil {
		return subscribers, session.TransactionError(ctx, err)
	}
	return subscribers, nil
}

func CountSubscribers(ctx context.Context, userId string) int {
	count, err := session.Database(ctx).Model((*Subscriber)(nil)).Where("user_id = ?", userId).Count()
	if err != nil {
		return 0
	}
	return count
}
