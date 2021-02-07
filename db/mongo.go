package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func InitMongoDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal(err)
	}
	os.Getenv("FOO")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("mongotest")

	//defer client.Disconnect(ctx)

	return db
}