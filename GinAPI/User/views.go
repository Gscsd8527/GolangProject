package User

import (
	result "GinAPI/Utils"
	"GinAPI/global"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

//注册用户
func RegisterView(c *gin.Context) result.Response{
	//db := DB.InitDB()
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	fmt.Println("username = ", username, "password= ", password, "email = ", email)
	if (username != "") && (password != "") && (email != ""){
		user := &Users{UserName: username, PassWord: password, Emali: email}
		var u1 Users
		//db.Where("username=? AND password=?", username, password).Find(&u1)
		global.DB.Where("username=? AND password=?", username, password).Find(&u1)
		// 根据用户名不相等才能注册成功
		if u1.UserName != username{
			//db.Create(&user)
			global.DB.Create(&user)

			//加入jwt校验， 登录成功就返回一个token
			token, _ := result.GenToken(username)
			data :=make(map[string]string, 2)
			data["token"] = token

			return result.ResponseSuccess("注册成功", data)
		}else {
			return result.ResponseError("用户名已经被注册", "")
		}
	}
	return result.ResponseError("用户名、密码、邮箱漏填", "")
}

//用户登录
func LoginView(c *gin.Context) result.Response{
	//db := DB.InitDB()
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("aaa = %T   len =%d  \n", password, len(password))
	fmt.Println("username = ", username, "password= ", password)
	if (username != "") && (password != ""){
		//user := &Users{UserName: username, PassWord: password}
		var u1 Users
		//db.Where("username=? AND password=?", username, password).Find(&u1)
		global.DB.Where("username=? AND password=?", username, password).Find(&u1)
		if u1.UserName == username && u1.PassWord == password{

			//加入jwt校验， 登录成功就返回一个token
			token, _ := result.GenToken(username)
			data :=make(map[string]string, 2)
			data["token"] = token

			return result.ResponseSuccess("登录成功", data)
		}
		return result.ResponseError("登录失败，用户名密码错误", "")
	}
	return result.ResponseError("用户名密码漏填", "")
}

