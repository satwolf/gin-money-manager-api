package service

import (
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/user/dto"
	"gin-money-manager-api/modules/user/entity"
	"gin-money-manager-api/modules/user/repository"
)

type UserCreatorService struct {
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
}

func NewUserCreatorService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
) UserCreatorService {
	return UserCreatorService{
		userRepository: userRepository,
		roleRepository: roleRepository,
	}
}

func (s *UserCreatorService) CreateUser(body *dto.CreateUserRequest) (*entity.User, error) {
	var roles []entity.Role

	s.roleRepository.DB().Where("id IN ?", body.Roles).Find(&roles)

	hashedPassword, err := helper.HashPassword(body.Password)

	if err != nil {
		return nil, err
	}

	newUser := entity.User{
		Name:     body.Name,
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPassword,
		Roles:    roles,
	}

	user, err := s.userRepository.Create(newUser)

	return &user, err
}
