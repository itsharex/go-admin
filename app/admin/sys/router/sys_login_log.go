package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/sys/apis"

	"go-admin/core/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysLoginLogRouter)
}

// 需认证的路由代码
func registerSysLoginLogRouter(v1 *gin.RouterGroup) {
	api := apis.SysLoginLog{}

	r := v1.Group("/admin/sys/sys-login-log").Use(middleware.Auth()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
		r.GET("/export", api.Export)
	}
}
