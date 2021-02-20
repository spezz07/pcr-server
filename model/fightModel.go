package model

import (
	"pcrweb/impl"
)

type FightResult struct {
	CreatorName   string `json:"creatorName"`
	CreatorId     string `json:"creatorId"`
	ApproveNum    int    `json:"approveNum"`
	DisapproveNum int    `json:"disapproveNum"`
	Offensive     string `json:"offensive"`
	Defensive     string `json:"defensive"`
	SortOffensive string
	SortDefensive string
	MappingString string `json:"mappingString"`
	Description   string `json:"description"`
	Title         string `json:"title"`
	//MappingName   string `json:"mappingName"`

	AreaType  int `json:"areaType"`
	RId       int `json:"rId"`
	IsDel     int `json:"isDel"`
	BaseModel `xorm:"extends"`
}

type FightResultPageList struct {
	CreatorName   string          `json:"creatorName"`
	CreatorId     string          `json:"creatorId"`
	ApproveNum    int             `json:"approveNum"`
	DisapproveNum int             `json:"disapproveNum"`
	Offensive     []int           `json:"offensive"`
	Defensive     []int           `json:"defensive"`
	Status        int             `json:"status"`
	RId           int             `json:"rId"`
	Created       impl.TimeFormat `json:"created"`
	AreaType      int             `json:"areaType"`
	Description   string          `json:"description"`
	Title         string          `json:"title"`
	ApproveStatus int             `json:"approveStatus" xorm:"-"`
}

type FightResultPageInfo struct {
	Offensive []int `json:"offensive"`
	Defensive []int `json:"defensive"`
	AreaType  int   `json:"areaType"`
	UserId    string
	Pagination
	//MappingName   string `json:"mappingName"`
}

type FightResultInfo struct {
	Offensive     []int  `json:"offensive"`
	Defensive     []int  `json:"defensive"`
	MappingString string `json:"mappingString"`
	AreaType      int    `json:"areaType"`
	Description   string `json:"description"`
	Title         string `json:"title"`
	UserId        string
	//MappingName   string `json:"mappingName"`
}
type FightLike struct {
	BaseModel     `xorm:"extends"`
	Status        int `json:"status"`
	FightResultId int `json:"fightResultId"`
	UserId        string
}

type FightApprove struct {
	ResultRelationId int
	ApproveUserId    string
	DisapproveUserId string
	BaseModel        `xorm:"extends"`
}
