package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seven1122/ginBlog/pkg/erorrs"
	"seven1122/ginBlog/pkg/logging"
	"seven1122/ginBlog/pkg/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = erorrs.SUCCESS
		token := c.Query("token")
		if token == ""{
			code = erorrs.INVALID_PARAMS
		}else {
			claims, err := utils.ParseToken(token)
			if err != nil{
				code = erorrs.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt{
				code = erorrs.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != erorrs.SUCCESS{
			logging.Info("token 已过期！请重新登录")
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": erorrs.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()

	}

}
