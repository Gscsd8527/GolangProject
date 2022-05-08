package jwt

import (
	"GinRouter/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


//校验JWT
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = utils.SUCCESS
		Authorization := c.GetHeader("Authorization")
		token := strings.Split(Authorization, " ")
		fmt.Println("Authorization = ",Authorization)
		//请求路径
		request_url := c.FullPath()
		fmt.Println("请求路径为 ", request_url)
		//过滤掉不需要使用jwt的路径
		if request_url == "/user/login"{
			return
		}

		if Authorization == ""{
			code = utils.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token[0])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = utils.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = utils.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != utils.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  utils.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}