package controllers

import (
	"gitcom.com/n4ze3m/go-boil/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) GetUsers(ctx *gin.Context) {
	var users []models.User
	cursor, err := c.database.Collection("user").Find(c.ctx, bson.M{})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err = cursor.All(c.ctx, &users)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Create()
	_, err = c.database.Collection("user").InsertOne(c.ctx, user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
