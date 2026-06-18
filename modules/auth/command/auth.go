package command

import (
	"gin-money-manager-api/modules/auth/dto"
	"gin-money-manager-api/modules/auth/service"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/entity"

	"github.com/gin-gonic/gin"
)

type UserFinder interface {
	FindByUsername(username string) (*entity.User, error)
}

type AuthCommand struct {
	userRepository UserFinder
}

func NewAuthCommand(repository UserFinder) *AuthCommand {
	return &AuthCommand{
		userRepository: repository,
	}
}

func (command *AuthCommand) Handler(
	c *gin.Context,
) {
	var body dto.LoginRequest

	if !helper.BindAndValidate(c, &body) {
		return
	}

	user, err := command.userRepository.FindByUsername(
		body.Username,
	)

	if err != nil || user == nil {
		response.ServerError(c, "user not found")
		return
	}

	if !helper.CheckPassword(
		body.Password,
		user.Password,
	) {
		response.ServerError(c, err.Error())
		return
	}

	token, err := service.GenerateToken(
		user.ID.String(),
		user.Username,
	)

	if err != nil {
		response.ServerError(c, err.Error())
		return
	}

	response.Success(
		c,
		map[string]string{
			"token": token,
		},
		err,
	)
}
