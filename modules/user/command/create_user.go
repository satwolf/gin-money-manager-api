package command

import (
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/response"
	"gin-money-manager-api/modules/user/dto"
	"gin-money-manager-api/modules/user/service"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	creator service.UserCreatorService
}

func NewCreateUser(
	creator *service.UserCreatorService,
) *CreateUser {
	return &CreateUser{
		creator: *creator,
	}
}

func (h *CreateUser) Handler(c *gin.Context) {
	var body dto.CreateUserRequest

	if !helper.BindAndValidate(c, &body) {
		return
	}

	user, err := h.creator.CreateUser(&body)

	response.Created(c, user, err)
}
