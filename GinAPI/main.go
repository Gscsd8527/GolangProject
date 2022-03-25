package main

import (
	user_route "GinAPI/User"
	result "GinAPI/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func init() {
//	DB.InitDB()   // 自动创建表
//}


func main() {
	r := gin.Default()
	//创建一个路由组，游客可以访问
	guest := r.Group("/user")
	{
		guest.GET("/ping", func(c *gin.Context) {

			resp := result.ResponseSuccess("成功", "hello world")
			fmt.Println("resp = ", resp)
			c.JSON(http.StatusOK, resp)
		})
		guest.POST("/login", func(c *gin.Context) {
			resp := user_route.LoginView(c)
			fmt.Println("resp = ", resp)
			c.JSON(http.StatusOK, resp)
		})
		guest.POST("/register", func(c *gin.Context) {
			resp := user_route.RegisterView(c)
			c.JSON(http.StatusOK, resp)
		})
	}

	//使用jwt
	viewsSet := r.Group("/view", result.JWTAuthMiddleware())
	{
		viewsSet.GET("/show", func(c *gin.Context) {
			resp := result.ResponseSuccess("成功", "hello world")
			fmt.Println("resp = ", resp)
			c.JSON(http.StatusOK, resp)
		})
	}

	r.Run(":8080")
}

