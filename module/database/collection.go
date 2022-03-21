package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	players = "players"
)

func PlayerCol() *mongo.Collection {
	return db.Collection(players)
}
