package container

import (
	"gin-money-manager-api/modules/user/repository"
	"gin-money-manager-api/modules/user/service"
)

type UserContainer struct {
	RoleRepository repository.RoleRepository
	UserRepository repository.UserRepository

	UserCreatorService service.UserCreatorService
}
