package auth

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"admin": "1234",
	})
}
