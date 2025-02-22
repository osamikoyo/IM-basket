package data

import (
	"github.com/osamikoyo/IM-basket/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) AddProduct(UserID uint64, product models.Product) error {
	update := bson.M{
		"$push": bson.M{
			"products" : product,
		},
	}

	filter := bson.M{
		"user_id" : UserID,
	}

	_, err := s.coll.UpdateOne(s.ctx, filter, update)
	return err
}

func (s *Storage) DeleteProduct(UserID, productID uint64) error {
	filter := bson.M{
		"user_id" : UserID,
	}

	update := bson.M{
		"$pull" : bson.M{
			"products" : bson.M{
				"id" : productID,
			},
		},
	}
	_, err := s.coll.UpdateOne(s.ctx, filter, update)
	return err
}