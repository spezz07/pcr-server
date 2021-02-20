package model

import (
	"pcrweb/impl"
)

type PcrRole struct {
	RoleName          string `json:"roleName"`
	RoleType          int    `json:"roleType"`
	RoleAlias         string `json:"roleAlias"`
	RoleDesc          string `json:"roleDesc"`
	RoleImgId         string `json:"roleImgId"`
	RoleImgName       string `json:"roleImgName"`
	RoleImgUrl        string `json:"roleImgUrl"`
	RoleStar          int    `json:"roleStar"`
	Iscn              int    `json:"isCN"`
	Istw              int    `json:"isTW"`
	Isjp              int    `json:"isJP"`
	UnitId            int    `json:"unitId"`
	Pos               int    `json:"pos"`
	Comment           string `json:"comment"`
	MoveSpeed         int    `json:"moveSpeed"`
	Kana              string `json:"kana"`
	PrefabId          int    `json:"prefabId"`
	IsLimited         int    `json:"isLimited"`
	Rarity            int    `json:"rarity"`
	MotionType        int    `json:"motionType"`
	SeType            int    `json:"seType"`
	AtkType           int    `json:"atkType"`
	NormalAtkCastTime int    `json:"normalAtkCastTime"`
	SearchAreaWidth   int    `json:"searchAreaWidth"`
	Cutin_1           int    `json:"cutin_1"`
	Cutin_2           int    `json:"cutin_2"`
	Cutin1Star6       int    `json:"cutin1_Star6"`
	Cutin2Star6       int    `json:"cutin2_Star6"`
	GuildId           int    `json:"guildId"`
	ExskillDisplay    int    `json:"exskillDisplay"`
	OnlyDispOwned     int    `json:"onlyDispOwned"`
	startTime         string `json:"startTime"`
	endTime           string `json:"endTime"`
	BaseModel         `xorm:"extends"`
}
type UnitData struct {
	UnitId            int     `json:"unitId"`
	Pos               int     `json:"pos" xorm:"-"`
	Comment           string  `json:"comment"`
	MoveSpeed         int     `json:"moveSpeed"`
	Kana              string  `json:"kana"`
	PrefabId          int     `json:"prefabId"`
	IsLimited         int     `json:"isLimited"`
	Rarity            int     `json:"rarity"`
	MotionType        int     `json:"motionType"`
	SeType            int     `json:"seType"`
	AtkType           int     `json:"atkType"`
	NormalAtkCastTime float32 `json:"normalAtkCastTime"`
	SearchAreaWidth   int     `json:"searchAreaWidth"`
	Cutin_1           int     `json:"cutin_1"`
	Cutin_2           int     `json:"cutin_2"`
	Cutin1Star6       int     `json:"cutin1_Star6"`
	Cutin2Star6       int     `json:"cutin2_Star6"`
	GuildId           int     `json:"guildId"`
	ExskillDisplay    int     `json:"exskillDisplay"`
	OnlyDispOwned     int     `json:"onlyDispOwned"`
	StartTime         string  `json:"startTime"`
	EndTime           string  `json:"endTime"`
}

type PcrRoleInfo struct {
	Id          int             `json:"id"`
	RoleName    string          `json:"roleName"`
	RoleType    int             `json:"roleType"`
	RoleAlias   string          `json:"roleAlias"`
	RoleDesc    string          `json:"roleDesc"`
	RoleImgId   string          `json:"roleImgId"`
	RoleImgName string          `json:"roleImgName"`
	RoleImgUrl  string          `json:"roleImgUrl"`
	RoleStar    int             `json:"roleStar"`
	Pos         int             `json:"pos"`
	UnitId      int             `json:"unitId"`
	Iscn        int             `json:"isCN"`
	Istw        int             `json:"isTW"`
	Isjp        int             `json:"isJP"`
	ImgCloudUrl string          `json:"imgCloudUrl"`
	Created     impl.TimeFormat `json:"created"`
	Updated     impl.TimeFormat `json:"updated"`
}
type PcrRolePageInfo struct { // 用户返回struct
	Id          int    `json:"id"`
	RoleName    string `json:"roleName"`
	RoleType    int    `json:"roleType"`
	RoleAlias   string `json:"roleAlias"`
	RoleDesc    string `json:"roleDesc"`
	RoleImgId   string `json:"roleImgId"`
	RoleImgName string `json:"roleImgName"`
	RoleImgUrl  string `json:"roleImgUrl"`
	RoleStar    int    `json:"roleStar"`
	ImgCloudUrl string `json:"imgCloudUrl"`
	Iscn        int    `json:"isCN"`
	Istw        int    `json:"isTW"`
	Isjp        int    `json:"isJP"`
	Pos         int    `json:"pos"`
	UnitId      int    `json:"unitId"`
}
type PcrRolePageList struct {
	PcrRoleInfo `xorm:"extends"`
	Pagination  `xorm:"extends"`
}
