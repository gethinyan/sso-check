package v1

import (
	"net/http"
	"time"

	"github.com/gethinyan/sso-check/models"
	"github.com/gethinyan/sso-check/pkg/setting"
	"github.com/gin-gonic/gin"
)

// LoginCheckRequest 用户注册请求参数
// swagger:parameters LoginCheckRequest
type LoginCheckRequest struct {
	// in: body
	Body LoginCheckRequestBody
}

// LoginCheckRequestBody 用户登录参数
type LoginCheckRequestBody struct {
	// 邮箱地址
	// Required: true
	Email string `json:"email"`
	// 密码
	// Required: true
	Password string `json:"password"`
}

// LoginCheckResponse 用户注册/登录响应参数
// swagger:response LoginCheckResponse
type LoginCheckResponse struct {
	// in: body
	Body struct {
		// 响应信息
		Message string `json:"message"`
		// 用户信息
		Data models.UserResponseBody `json:"data"`
	}
}


// LoginCheck swagger:route POST /loginCheck LoginCheckRequest
//
// 登录检查
//
//     Schemes: http, https
//
//     Responses:
//       200: LoginCheckResponse
func LoginCheck(c *gin.Context) {
	c.JSON(http.StatusOK, c.user)
}
