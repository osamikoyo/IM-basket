package data

import (
	"github.com/osamikoyo/IM-basket/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Storage) Create(userID uint64) error {
	var basket models.Basket
	basket.UserID = userID

	_, err := s.coll.InsertOne(s.ctx, basket)
	return err
}

func (s *Storage) Get(userID uint64) (*models.Basket, error) {
	cursor :=  s.coll.FindOne(s.ctx, bson.M{
		"user_id" : userID,
	})

	var basket models.Basket
	err := cursor.Decode(&basket)
	return &basket, err
}