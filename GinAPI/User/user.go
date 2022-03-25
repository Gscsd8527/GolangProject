package User

type Users struct {
	Id	int `xorm:"pk autoincr" gorm:"pk;auto"`
	UserName string `json:"username" gorm:"column:username;size:255"`
	PassWord string `json:"password" gorm:"column:password;size:255"`
	Emali string `json:"email" gorm:"email"`
}

