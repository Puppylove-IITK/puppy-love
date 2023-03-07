package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PuppyDb struct {
	Session *mongo.Client
}

func MongoConnect() (*PuppyDb, error) {
	// Set your MongoDB Atlas connection string here.
	connectionString := "mongodb+srv://username:securepassword@puppylove.woq42jd.mongodb.net/?retryWrites=true&w=majority"

	// Set client options.
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB Atlas cluster.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB Atlas: %v", err)
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB Atlas: %v", err)
	}

	fmt.Println("Connected to MongoDB Atlas successfully!")

	return &PuppyDb{Session: client}, nil
}

func (db *PuppyDb) GetCollection(collection string) *mongo.Collection {
	return db.Session.Database("puppy").Collection(collection)
}

func (db *PuppyDb) GetById(collection string, id interface{}) (*mongo.SingleResult, error) {
	c := db.GetCollection(collection)
	return c.FindOne(context.Background(), id), nil
}
