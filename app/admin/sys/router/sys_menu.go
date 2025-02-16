package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/apis"

	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysMenuRouter)
}

// 需认证的路由代码
func registerSysMenuRouter(v1 *gin.RouterGroup) {
	api := apis.SysMenu{}

	r := v1.Group("/admin/sys/sys-menu").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetTreeList)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/menu-role", api.GetMenuRole)
		r.GET("/role-menu-tree-select/:roleId", api.GetMenuTreeSelect)
	}
}
