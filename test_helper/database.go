package test_helper

import (
	"MongodbTest/config"
	"MongodbTest/model"
	"MongodbTest/module/database"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// CreateFakePlayer ...
func CreateFakePlayer() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	var (
		playerCol  = database.PlayerCol()
		ctx        = context.Background()
		fakePlayer = model.Player{
			ID:   primitive.NewObjectID(),
			Name: "fake", Email: "fake@gmail.com",
			Password:  string(bytes),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	)

	//Insert
	_, err := playerCol.InsertOne(ctx, fakePlayer)
	if err != nil {
		log.Println(err)
	}
}

// ConnectTestDB ...
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
