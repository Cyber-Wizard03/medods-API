package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {

	router := gin.Default()

	router.Use(cors.Default())
	api := router.Group("/api")
	{

		api.GET("/FeedbackMessages", FeedbackGet)
		api.POST("/FeedbackMessages", FeedbackPost)

	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
