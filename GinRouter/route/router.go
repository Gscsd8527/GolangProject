package route

import (
	"GinRouter/api"
	"GinRouter/middleware/jwt"
	"GinRouter/middleware/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


//初始化路由
func InitRoute() *gin.Engine {
	//r := gin.Default()

	// 创建一个不包含中间件的路由器
	r := gin.New()

	//r.Use(gin.Logger()) // 日志
	r.Use(logger.LoggerToFile()) // 日志
	r.Use(Cors())       // 跨域请求
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

	//首页路由
	r.GET("/",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//用户模块
	userRoute := r.Group("/user")

	userRoute.Use(jwt.JWT()) //加入token验证

	{
		userRoute.POST("/login", api.UserLogin)
		userRoute.POST("/register", api.UserRegister)
		userRoute.GET("/info", api.UserInfo)
	}

	return r
}


// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
