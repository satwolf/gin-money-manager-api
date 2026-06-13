package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/user/command"
	"gin-money-manager-api/modules/user/container"
	"gin-money-manager-api/modules/user/resource"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoute(r *gin.Engine, container *container.UserContainer) {
	createUser := command.NewCreateUser(&container.UserCreatorService)
	userResource := resource.NewUserResource(container.UserRepository)

	users := r.Group("/users")
	users.POST("", createUser.Handler)

	users.Use(
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("user"),
	)
	{
		users.GET("", userResource.Index)
	}
}
