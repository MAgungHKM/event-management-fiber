package tag

import (
	"event-management/model"
)

type TagResponse struct {
	model.ResponseWithData
	Data model.Tags `json:"data"`
}

type TagRequest struct {
	Name string `json:"name" validate:"required,alpha,max=100"`
}

func (request TagRequest) MapToTag(user *model.Tag) {
	user.Name = request.Name
}
