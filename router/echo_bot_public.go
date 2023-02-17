package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sam80719/echoBot/handler"
	"github.com/sam80719/echoBot/middleware"
)

func SetRouterPublic() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Auth(), gin.Recovery())

	userRouter := router.Group("/message")
	{
		userRouter.POST("/:id", handler.SaveMessage)
	}

	return router
}
