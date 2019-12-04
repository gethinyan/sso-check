package v1

import (
	"net/http"

	"github.com/gethinyan/sso-check/rpcx"
	"github.com/gin-gonic/gin"
)

// LoginCheck swagger:route POST /loginCheck LoginCheckRequest
//
// 登录检查
//
//     Schemes: http, https
//
//     Responses:
//       200: LoginCheckResponse
func LoginCheck(c *gin.Context) {
	c.JSON(http.StatusOK, rpcx.Reply)
}
