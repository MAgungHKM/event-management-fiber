package user

import (
	"event-management/model"
	"event-management/utils/secret"
)

type UserResponse struct {
	model.ResponseWithData
	Data model.Users `json:"data"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,ascii"`
}

func (request UserRequest) MapToUser(user *model.User) {
	user.Name = request.Name
	user.Email = request.Email
	user.Username = request.Username
	user.Secret = secret.GenerateHash(request.Password)
}
