package dto

import (
	"go-admin/core/dto"
)

type SysDictTypeQueryReq struct {
	dto.Pagination `search:"-"`
	DictName       string `form:"dictName" search:"type:icontains;column:dict_name;table:admin_sys_dict_type"`
	DictType       string `form:"dictType" search:"type:icontains;column:dict_type;table:admin_sys_dict_type"`
	BeginCreatedAt string `form:"beginCreatedAt" search:"type:gte;column:created_at;table:admin_sys_dict_type" comment:"创建时间"`
	EndCreatedAt   string `form:"endCreatedAt" search:"type:lte;column:created_at;table:admin_sys_dict_type" comment:"创建时间"`
}

type SysDictTypeOrder struct {
	IdOrder string `form:"idOrder" search:"type:order;column:id;table:admin_sys_dict_type"`
}

func (m *SysDictTypeQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysDictTypeInsertReq struct {
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Remark     string `json:"remark"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysDictTypeUpdateReq struct {
	Id         int64  `uri:"id" json:"-"`
	DictName   string `json:"dictName"`
	DictType   string `json:"dictType"`
	Remark     string `json:"remark"`
	CurrUserId int64  `json:"-" comment:""`
}

type SysDictTypeGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

// SysDictDeleteReq 功能删除请求参数
type SysDictrDeleteReq struct {
	Ids []int64 `json:"ids"`
}
