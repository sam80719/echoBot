package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	// handle user
	return func(c *gin.Context) {
		println("auth test \n")
		c.Next()
	}
}
