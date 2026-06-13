package routes

import (
	"gin-money-manager-api/modules/user/container"
	"gin-money-manager-api/modules/user/resource"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoute(r *gin.Engine, container *container.UserContainer) {
	roleResource := resource.NewRoleResource(container.RoleRepository)

	roles := r.Group("/roles")
	{
		roles.GET("", roleResource.Index)
		roles.GET("/:id", roleResource.Show)
	}
}
