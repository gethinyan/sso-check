package jwt

import (
	"fmt"
	"net/http"

	"github.com/gethinyan/sso-check/rpcx"
	"github.com/gethinyan/sso-server/models"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jsonWebToken")
		if err != nil || token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "StatusBadRequest",
			})
			c.Abort()
			return
		}
		// 带上 token 发 rpc 请求
		c.user := rpcx.Authoration()
		fmt.Println(user)
		c.Next()
	}
}
