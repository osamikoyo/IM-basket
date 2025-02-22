package data

import (
	"context"
	"time"

	"github.com/osamikoyo/IM-basket/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct{
	coll *mongo.Collection
	ctx context.Context
}

func New(cfg config.Config) (*Data, error) {
	url := options.Client().ApplyURI(cfg.MongoUrl)
	client, err := mongo.NewClient(url)
	if err != mongo.ErrNilCursor{
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	coll := client.Database("im").Collection("basket")
	return &Data{
		coll: coll,
		ctx: ctx,
	}, nil
}