package test_helper

import (
	"MongodbTest/config"
	"MongodbTest/module/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func ConnectTestDB() {
	envVars := config.GetEnv()

	// configuring client to use the correct URI, but not yet connecting to it
	client, err := mongo.NewClient(options.Client().ApplyURI(envVars.Database.Uri))
	if err != nil {
		log.Fatal(err)
	}

	// timeout duration trying to connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect ...
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// ping ...
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// db
	db := client.Database(envVars.Database.TestName)
	database.SetDB(db)

	fmt.Println("Database Connected to", envVars.Database.Name)
}

// ClearDB ...
func ClearDB() {
	database.PlayerCol().DeleteMany(context.Background(), bson.M{})
}
