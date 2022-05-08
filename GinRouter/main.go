package main

import (
	"GinRouter/config"
	"GinRouter/database"
	"GinRouter/pkg/setting"
	"GinRouter/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main() {
	// 注册路由
	r := route.InitRoute()

	// 加载配置文件
	config.InitConfig()

	// 创建数据库连接
	database.InitDB()

	//默认http服务器启动服务
	//gin.SetMode(viper.GetString("server.run_mode"))
	//r.Run(":" + viper.GetString("server.port"))

	//自定义http服务器
	gin.SetMode(setting.ServerSetting.RunMode)
	r.Run(":" + fmt.Sprintf("%d", setting.ServerSetting.HttpPort))

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()
}
