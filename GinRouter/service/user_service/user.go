package user_service

import (
	"GinRouter/models"
	"GinRouter/utils"
	"fmt"
	"reflect"
	"strings"
)


//登录用户所需要的字段
type LoginUserInfo struct {
	UserName string `json:"username" binding:"required" form:"username"`
	PassWord string  `json:"password" binding:"required" form:"password"`
}


//注册用户所需要的字段
type RegisterUserInfo struct {
	UserName string `json:"username" binding:"required" form:"username"`
	PassWord string  `json:"password" binding:"required" form:"password"`
	Email string  `json:"email" binding:"required" form:"email"`
}


//定义查看用户的字段信息
type LookUserInfo struct {
	Id	int `json:"id"`
	UserName string `json:"username" `
	PassWord string  `json:"password"`
	Email string  `json:"email"`
}


//获取用户
func (userinfo *LoginUserInfo) UserGet() (*models.Users, error){
	fmt.Println("获取用户： ", userinfo)
	user, err := models.GetUser(userinfo.UserName, userinfo.PassWord)
	if err != nil {
		return nil, err
	}
	return user, nil
}


//注册用户
func (register_user *RegisterUserInfo) UserAdd() (*models.Users, string) {
	fmt.Println("注册用户： ", register_user)
	// 先去查询一下该用户名密码是否注册过
	count, err := register_user.Count()
	fmt.Println("COUNT = ", count)
	if (err != nil) || (count > 0){
		return nil, utils.GetMsg(utils.ERROR_USER_EXIST_FAIL)
	}

	//新增用户
	fmt.Println("新增用户")
	user, _ := models.AddUser(register_user.UserName, register_user.PassWord, register_user.Email)
	fmt.Println("用户是否注册成功", user)
	if user == nil {
		return nil, utils.GetMsg(utils.ERROR_USER_EXIST_FAIL)
	}
	return user, ""
}


//查看用户
func (userinfo *LookUserInfo) UserLook() (*models.Users, error){
	fmt.Println("查看用户信息： ", userinfo)
	user, err := models.GetUserId(userinfo.Id)
	if err != nil{
		fmt.Println("查看用户失败")
		return nil, err
	}
	return user, err
}


//统计数量
func (register_user *RegisterUserInfo) Count() (int64, error){
	return models.GetUserTotal(register_user.getMaps())
}


//组合条件: 不能写 *RegisterUserInfo，这个传地址
func (register_user RegisterUserInfo) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	//maps["username"] = register_user.UserName
	//maps["password"] = register_user.PassWord
	regiterType := reflect.TypeOf(register_user)
	regiterValue := reflect.ValueOf(register_user)
	fmt.Println(regiterType, regiterValue)
	for i := 0; i < regiterType.NumField(); i++ {
		tagName := regiterType.Field(i).Name
		tags := strings.Split(string(regiterType.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		value := regiterValue.Field(i).Interface()
		fmt.Println("tagName = ", tagName, value)

		//指定要的字段： 将对象里面的数据放入到map里面
		if tagName == "username" || tagName == "password"{
			maps[tagName] = value
		}
	}
	fmt.Println("组合条件 ->", maps)
	return maps
}

