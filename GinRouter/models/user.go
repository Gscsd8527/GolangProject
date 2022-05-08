package models

import (
	"GinRouter/global"
	"fmt"
	"gorm.io/gorm"
)

// 注册users表时的字段：Model类型
type Users struct {
	Id	int `xorm:"pk autoincr" gorm:"pk;auto"`
	UserName string `json:"username" gorm:"column:username;size:255"`
	PassWord string `json:"password" gorm:"column:password;size:255"`
	Emali string `json:"email" gorm:"email"`
}


//查看用户
func GetUser(username string, password string) (*Users, error){
	var user Users
	err := global.DB.Where("username=? AND password=?", username, password).First(&user).Error
	if err != nil{
		fmt.Println("查看用户失败")
		return nil, err
	}
	fmt.Println("查看用户成功", user)
	return &user, nil
}


//增加用户
func AddUser(username string, password string, email string) (*Users, error){
	user := &Users{
		UserName: username,
		PassWord: password,
		Emali: email,
	}
	// 创建用户
	err := global.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}


//统计用户数量
func GetUserTotal(maps interface{}) (int64, error){
	var count int64
	err := global.DB.Model(&Users{}).Where(maps).Count(&count).Error
	if err != nil{
		return 0, err
	}
	return count, nil
}


//获取单个用户
func GetUserId(id int)(*Users, error){
	var user Users
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	fmt.Println("获取单个用户 = ", user)
	return &user, nil
}