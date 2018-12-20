package models

import (
	"context"
	"github.com/MixinNetwork/bot-api-go-client"
	"github.com/MixinNetwork/go-number"
	"github.com/caosbad/ever-post-mixin-bot/session"
	"github.com/satori/go.uuid"
	"time"
)

const supports_DDL = `
CREATE TABLE IF NOT EXISTS supports (
  support_id            	VARCHAR(36) PRIMARY KEY,
  supporter_id	        	VARCHAR(36) NOT NULL,
  author_id					VARCHAR(36) NOT NULL,
  post_id					VARCHAR(36) NOT NULL,
  asset_id         			VARCHAR(36) NOT NULL,
  trace_id         			VARCHAR(36) NOT NULL,
  amount        			VARCHAR(256),
  created_at        	TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
);

CREATE INDEX ON posts (support_id);
`

type Support struct {
	SupportId   string         `sql:"support_id,pk"`
	SupporterId string         `sql:"user_id,notnull"`
	AuthorId    string         `sql:"author_id,notnull"`
	AssetId     string         `sql:"asset_id,notnull"`
	PostId      string         `sql:"post_id,notnull"`
	TraceId     string         `sql:"trace_id,notnull"`
	Amount      number.Decimal `sql:"amount,notnull"`
	CreatedAt   time.Time      `sql:"created_at,notnull"`
}

func CreateSupport(ctx context.Context, post *Post, transfer bot.TransferView) (*Support, error) {
	support := &Support{
		SupportId:   uuid.Must(uuid.NewV4()).String(),
		SupporterId: transfer.CounterUserId,
		AuthorId:    post.UserId,
		AssetId:     transfer.AssetId,
		PostId:      post.PostId,
		TraceId:     transfer.TraceId,
		Amount:      number.FromString(transfer.Amount),
		CreatedAt:   time.Now(),
	}
	if err := session.Database(ctx).Insert(support); err != nil {
		return nil, session.TransactionError(ctx, err)
	}
	return support, nil
}
