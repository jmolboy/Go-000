package dao

type User struct {
	AppModel
	OpenId    string `gorm:"column:openid" json:"openid"`
	NickName  string `gorm:"column:nickname" json:"nickname"`
	AvatarUrl string `gorm:"column:avatarurl" json:"avatarurl"`
}

func (User) TableName() string {
	return "st_user"
}

/**
 添加
 */
func (user *User) Create() (insertId int, err error) {
	MyDB.Create(user)
	insertId = user.Id
	return
}

