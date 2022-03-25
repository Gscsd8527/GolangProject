package DB

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB {
	//前提是你要先在本机用Navicat创建一个名为go_db的数据库
	host := "127.0.0.1"
	port := "3306"
	database := "gin_demo"
	username := "root"
	password := "bigdata"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	//这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	//这个是gorm自动创建数据表的函数。它会自动在数据库中创建一个名为users的数据表
	//_ = db.AutoMigrate(&User.Users{}) // 创建后可注销掉
	return db
}
