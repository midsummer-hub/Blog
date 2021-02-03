package entity


//博主
type Blogger struct {
	Id 		 int 	`gorm:"column:id"json:"id"`					//ID
	Username string `gorm:"column:username"json:"username"`		//用户名
	Password string `gorm:"column:password"json:"password"`		//密码
	Nickname string `gorm:"column:nickname"json:"nickname"`		//昵称
	Sign	 string `gorm:"column:sign"json:"sign"`				//个人签名
	Profile  string `gorm:"column:profile"json:"profile"`		//个人简介
	Img 	 string `gorm:"column:img"json:"img"`				//个人头像
}
