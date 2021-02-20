package model

import "time"

type Pagination struct {
	Rows   int   `json:"rows"`
	PageNo int   `json:"pageNo"`
	Total  int64 `json:"total"`
}

type BaseModel struct {
	Id      int64     `xorm:"pk"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type BaseRequest interface {
	GetPageList(params interface{})
	GetList(params interface{})
}
