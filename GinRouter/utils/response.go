package utils

import "github.com/gin-gonic/gin"


type Gin struct {
	C *gin.Context
}


//定义返回值
func (g *Gin) Response(httpCode, Code int, data interface{}){
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg": GetMsg(Code),
		"data": data,
	})
	return
}