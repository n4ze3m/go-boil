package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInit(ctx context.Context) (*mongo.Client, error) {
	uri := os.Getenv("DATABASE_URL")
	return mongo.Connect(ctx, options.Client().ApplyURI(uri))
}
