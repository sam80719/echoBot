package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sam80719/echoBot/handler"
	"github.com/sam80719/echoBot/middleware"
	"net/http"
)

func SetRouterPublic() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Auth(), gin.Recovery())

	userRouter := router.Group("/api")
	{

		userRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  200001,
				"message": "health check ok",
			})
		})
		userRouter.POST("/message", handler.SaveMessage)
	}

	return router
}
