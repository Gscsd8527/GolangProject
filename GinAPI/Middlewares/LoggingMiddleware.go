package Middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}


// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := "./"
	logFileName := "gin.log"

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志  文本  格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	//设置日志  json  格式
	//logger.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		//响应的数据
		response_data := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = response_data

		// 处理请求
		// 请求前
		c.Next()
		// 请求后


		// 接收响应的数据并装换成json格式
		var f interface{}
		json.Unmarshal(response_data.body.Bytes(), &f)
		m := f.(map[string]interface{})
		fmt.Println("接口真实响应的数据为： ", m)
		fmt.Println("取值code : ", m["code"])


		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()  //接口的状态码
		//statusCode :=  m["code"] //取响应数据里面的值

		// 请求IP
		clientIP := c.ClientIP()


		// 日志格式
		//logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//	statusCode,
		//	latencyTime,
		//	clientIP,
		//	reqMethod,
		//	reqUri,
		//)


		//转换成json
		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
			"response_data": response_data.body.String(),
		}).Info()

	}
}


// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}