package server

import (
	"context"
	"log"

	"gitcom.com/n4ze3m/go-boil/controllers"
	"gitcom.com/n4ze3m/go-boil/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ctx context.Context
var controller *controllers.Controller

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx = context.Background()
	client, err := database.MongoInit(ctx)
	if err != nil {
		log.Fatalf("MongoInit error: %v", err)
	}
	controller = controllers.NewController(ctx, client.Database("mock"))
}

func routerInit() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	{
		v1 := api.Group("/v1")

		{
			user := v1.Group("/user")
			{
				user.GET("/", controller.GetUsers)
				user.POST("/", controller.CreateUser)
			}
		}
	}

	return router
}
