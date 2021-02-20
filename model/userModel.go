package model

import "pcrweb/impl"

type UserModel struct {
	BaseModel  `xorm:"extends"`
	Password   string `json:"password" xorm:"column:password `
	UserId     string `json:"userId" gorm:"xorm:user_id;NOT NULL;unique" `
	UserName   string `json:"userName" xorm:"user_name"`
	RoleId     int    `json:"roleId" gorm:"xorm:role_id"`
	Account    string `json:"account" xorm:"account"`
	OpenId     string `json:"openId" xorm:"open_id"`
	Avatar     string `json:"avatar"`
	UserStatus int    `json:"userStatus"`
	Permission string `json:"permission"`
	UserType   int
}

type UserWxModel struct {
	BaseModel  `xorm:"extends"`
	Password   string `json:"password" xorm:"column:password`
	UserId     string `json:"userId" gorm:"xorm:user_id;NOT NULL;unique" `
	UserName   string `json:"userName" xorm:"user_name"`
	RoleId     int    `json:"roleId" gorm:"xorm:role_id"`
	Account    string `json:"account" xorm:"account"`
	OpenId     string `json:"openId" xorm:"open_id"`
	Avatar     string `json:"avatar"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Province   string `json:"province"`
	Gender     int    `json:"gender"`
	UserStatus int    `json:"userStatus"`
	Permission string `json:"permission"`
	UserType   int
	SessionKey string
	Code       string `json:"code" xorm:"-"`
}

type UserInfo struct {
	UserId   string          `json:"userId" gorm:"xorm:user_id;NOT NULL;unique" `
	UserName string          `json:"userName" xorm:"user_name"`
	Avatar   string          `json:"avatar"`
	Account  string          `json:"account"`
	Created  impl.TimeFormat `json:"created"`
}
type UserList struct {
	Pagination `xorm:"-"`
	UserInfo
}

func (UserModel) TableName() string {
	return "pcr_user"
}

func (UserList) TableName() string {
	return "pcr_user"
}
