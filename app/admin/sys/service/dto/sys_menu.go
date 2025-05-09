package dto

import (
	"go-admin/app/admin/sys/models"
	"go-admin/core/dto"
)

// SysMenuQueryReq 列表或者搜索使用结构体
type SysMenuQueryReq struct {
	dto.Pagination `search:"-"`
	Id             string  `form:"id" search:"type:exact;column:id;table:admin_sys_menu" comment:"菜单编号"`                        // 菜单编号
	Title          string  `form:"title" search:"type:contains;column:title;table:admin_sys_menu" comment:"菜单名称"`               // 菜单名称
	Path           string  `form:"path" search:"type:exact;column:path;table:admin_sys_menu" comment:"路由地址"`                    // 路由地址
	IsHidden       string  `form:"isHidden" search:"type:exact;column:is_hidden;table:admin_sys_menu" comment:"显示状态 1-隐藏 2-显示"` // 显示状态
	ParentId       int64   `form:"-" search:"type:exact;column:parent_id;table:admin_sys_menu" comment:"父级"`                    // 父级
	ParentIds      []int64 `form:"-" search:"type:in;column:parent_id;table:admin_sys_menu" comment:"父级"`                       // 父级
	MenuIds        []int64 `form:"-" search:"type:in;column:id;table:admin_sys_menu" comment:"菜单编号"`                            // 菜单编号
}

func (m *SysMenuQueryReq) GetNeedSearch() interface{} {
	return *m
}

type SysMenuInsertReq struct {
	Title       string          `form:"title" comment:"显示名称"`  //显示名称
	Icon        string          `form:"icon" comment:"图标"`     //图标
	Path        string          `form:"path" comment:"路径"`     //路径
	Redirect    string          `form:"redirect" comment:"跳转"` //针对目录跳转，比如搜索出菜单
	Element     string          `form:"element" comment:"组件"`  //组件
	SysApi      []models.SysApi `form:"sysApi"`
	Apis        []int           `form:"apis"`
	Permission  string          `form:"permission" comment:"权限编码"` //权限编码
	ParentId    int64           `form:"parentId" comment:"上级菜单"`   //上级菜单
	Sort        int             `form:"sort" comment:"排序"`
	MenuType    string          `form:"menuType" comment:"菜单类型"`            //菜单类型
	IsKeepAlive string          `form:"isKeepAlive" comment:"是否缓存 1-是 2-否"` //是否缓存
	IsAffix     string          `form:"isAffix" comment:"是否固定 1-是 2-否"`     //是否固定
	IsHidden    string          `form:"isHidden" comment:"1-隐藏 2-显示"`       //是否显示
	IsFrame     string          `form:"isFrame" comment:"内嵌 1-是 2-否"`       //是否frame
	CurrUserId  int64           `json:"-" comment:""`
}

type SysMenuUpdateReq struct {
	Id          int64           `uri:"id" json:"-" comment:"编码"` // 编码
	Title       string          `form:"title" comment:"显示名称"`    //显示名称
	Icon        string          `form:"icon" comment:"图标"`       //图标
	Path        string          `form:"path" comment:"路径"`       //路径
	Redirect    string          `form:"redirect" comment:"跳转"`   //针对目录跳转，比如搜索出菜单
	Element     string          `form:"element" comment:"组件"`    //组件
	SysApi      []models.SysApi `form:"sysApi"`
	Apis        []int           `form:"apis"`
	Permission  string          `form:"permission" comment:"权限编码"`          //权限编码
	ParentId    int64           `form:"parentId" comment:"上级菜单"`            //上级菜单
	Sort        int             `form:"sort" comment:"排序"`                  //排序
	MenuType    string          `form:"menuType" comment:"菜单类型"`            //菜单类型
	IsKeepAlive string          `form:"isKeepAlive" comment:"是否缓存 1-是 2-否"` //是否缓存
	IsAffix     string          `form:"isAffix" comment:"是否固定 1-是 2-否"`     //是否固定
	IsHidden    string          `form:"isHidden" comment:"1-隐藏 2-显示"`       //是否显示
	IsFrame     string          `form:"isFrame" comment:"内嵌 1-是 2-否"`       //是否frame
	CurrUserId  int64           `json:"-" comment:""`
}

type SysMenuDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type SysMenuGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

type SelectMenuRole struct {
	RoleId int64 `uri:"roleId"`
}

type MenuTreeRoleResp struct {
	Menus       []*models.SysMenu `json:"menus"`
	CheckedKeys []int64           `json:"checkedKeys"`
}
