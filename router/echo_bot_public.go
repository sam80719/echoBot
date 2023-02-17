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

	apiRouter := router.Group("/api")
	{

		apiRouter.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  200001,
				"message": "health check ok",
			})
		})

		apiRouter.POST("/callback", handler.HandleMessage)
	}

	return router
}
