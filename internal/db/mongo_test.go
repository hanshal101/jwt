package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// setup initializes the MongoDB connection and returns a reference to the collection.
func setup(t *testing.T) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	assert.NoError(t, err)

	err = client.Ping(context.Background(), nil)
	assert.NoError(t, err)

	db := client.Database("Users")
	collection := db.Collection("tests")

	return collection
}

func TestConnectionbyInsert(t *testing.T) {
	coll := setup(t)
	defer coll.Drop(context.Background())
	doc := bson.D{
		{Key: "key", Value: "value"},
	}
	_, err := coll.InsertOne(context.Background(), doc)
	assert.NoError(t, err)

	var result bson.M
	err = coll.FindOne(context.Background(), bson.D{}).Decode(&result)
	assert.NoError(t, err)

	assert.Equal(t, "value", result["key"])
}
