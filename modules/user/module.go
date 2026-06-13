package user

import (
	"gin-money-manager-api/modules/user/container"
	"gin-money-manager-api/modules/user/repository"
	"gin-money-manager-api/modules/user/routes"
	"gin-money-manager-api/modules/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)

	userCreatorService := service.NewUserCreatorService(userRepository, roleRepository)

	userContainer := &container.UserContainer{
		RoleRepository: roleRepository,
		UserRepository: userRepository,

		UserCreatorService: userCreatorService,
	}

	routes.RegisterRoleRoute(r, userContainer)
	routes.RegisterUserRoute(r, userContainer)
}
