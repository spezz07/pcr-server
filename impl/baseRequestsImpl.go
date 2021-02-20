package impl

import "pcrweb/response"

type BaseRequestsImpl interface {
	Add(v interface{}) (data interface{}, err error)
	Update(v interface{}) (data interface{}, err error)
	Delete(v interface{}) (data interface{}, err error)
	GetPageList(v interface{}) (page response.Pagination, data interface{}, err error)
	GetList(v interface{}) (data interface{}, err error)
}
