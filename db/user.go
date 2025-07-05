package db

import (
	"HotelReservation/types"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	dbname     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(dbNAME).Collection(userCollection),
	}
}

func (m *MongoUserStore) GetUserByID(ctx context.Context,
	id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := m.collection.FindOne(ctx, oid).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
