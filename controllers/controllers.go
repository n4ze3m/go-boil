package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	ctx   context.Context
	database *mongo.Database
}

func NewController(ctx context.Context, database *mongo.Database) *Controller {
	return &Controller{
		ctx: ctx,
		database: database,
	}
}