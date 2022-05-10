package api

import (
	"GinRouter/service/user_service"
	"GinRouter/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)


// @title 用户登录接口
// @version 1.0
// @description 用户登录
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /user/login  [POST]
func UserLogin(c *gin.Context){
	//将 Context 放入到 结构体Gin中，进行封装
	appG := utils.Gin{C:c}

	//两种方式都可：var userinfo models.LoginUserInfo
	userinfo := user_service.LoginUserInfo{}

	// ShouldBind： 获取请求结构体的内容
	err := c.ShouldBind(&userinfo)

	if err != nil{
		appG.Response(http.StatusOK, utils.INVALID_PARAMS, nil)
		return
	}

	user, err := userinfo.UserGet()
	if err != nil{
		appG.Response(http.StatusOK, utils.ERROR_GET_USER_FAIL, nil)
		return
	}

	//加入md5
	token, err := utils.GenerateToken(user.Id, user.UserName, user.PassWord)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR_AUTH_TOKEN, nil)
		return
	}

	//appG.Response(http.StatusOK, utils.SUCCESS, user)
	appG.Response(http.StatusOK, utils.SUCCESS, map[string]string{
		"token": token,
	})
	return
}


// @title 用户注册接口
// @version 1.0
// @description 用户注册
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /user/register  [POST]
func UserRegister(c *gin.Context){
	//将 Context 放入到 结构体Gin中，进行封装
	appG := utils.Gin{C:c}

	register_user := user_service.RegisterUserInfo{}

	// ShouldBind： 获取请求结构体的内容
	err := c.ShouldBind(&register_user)

	if err != nil{
		appG.Response(http.StatusOK, utils.INVALID_PARAMS, nil)
		return
	}

	user, _ := register_user.UserAdd()
	if user == nil{
		appG.Response(http.StatusOK, utils.ERROR_USER_EXIST_FAIL, nil)
		return
	}

	//加入md5
	token, err := utils.GenerateToken(user.Id, user.UserName, user.PassWord)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR_AUTH_TOKEN, nil)
		return
	}

	//appG.Response(http.StatusOK, utils.SUCCESS, user)
	//返回token值
	appG.Response(http.StatusOK, utils.SUCCESS, map[string]string{
		"token": token,
	})

	return
}


// @title 查看用户信息
// @version 1.0
// @description 用户信息
// @Accept json
// @Produce  json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /user/info  [GET]
func UserInfo(c *gin.Context){
	//将 Context 放入到 结构体Gin中，进行封装
	appG := utils.Gin{C:c}
	Authorization := c.GetHeader("Authorization")
	token := strings.Split(Authorization, " ")

	t, err := jwt.Parse(token[0], func(*jwt.Token) (interface{}, error) {
		return utils.JwtSecret, nil
	})
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR_AUTH, nil)
		return
	}
	fmt.Println("T = ", t)
	//获取的值
	user_id := utils.GetIdFromClaims("id", t.Claims)
	user_id_int, err := strconv.Atoi(user_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, utils.ERROR_AUTH, nil)
		return
	}

	look_userinfo := user_service.LookUserInfo{Id: user_id_int}
	user, err := look_userinfo.UserLook()
	if err != nil{
		appG.Response(http.StatusInternalServerError, utils.ERROR_AUTH, nil)
		return
	}

	appG.Response(http.StatusOK, utils.SUCCESS, user)
	return
}