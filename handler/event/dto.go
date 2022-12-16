package event

import (
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"

	"github.com/gofiber/fiber/v2"
)

type EventResponse struct {
	model.ResponseWithData
	Data model.Events `json:"data"`
}

type EventRequest struct {
	Name        string   `json:"name" validate:"required,ascii"`
	StartDate   string   `json:"start_date" validate:"required,ISO8601datetime" example:"YYYY-MM-DDTHH:mm:ss.SSSZ"`
	EndDate     string   `json:"end_date" validate:"required,ISO8601datetime" example:"YYYY-MM-DDTHH:mm:ss.SSSZ"`
	Description string   `json:"description" validate:"required,ascii"`
	Location    string   `json:"location" validate:"required,ascii"`
	Tags        []string `json:"tags" validate:"dive,alpha,max=100"`
}

func (request EventRequest) MapToEvent(event *model.Event) error {
	startDate, err := utils.ParseDateTimeInJakarta(request.StartDate)
	if err != nil {
		return err
	}

	endDate, err := utils.ParseDateTimeInJakarta(request.EndDate)
	if err != nil {
		return err
	}

	if !startDate.Before(*endDate) {
		return errors.New("start_date must be lesser than end_date", fiber.StatusBadRequest)
	}

	event.Name = request.Name
	event.StartDate = *startDate
	event.EndDate = *endDate
	event.Description = request.Description
	event.Location = request.Location

	return nil
}

type EventParticipantRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (request EventParticipantRequest) MapToEvent(eventParticipant *model.EventParticipant) {
	eventParticipant.Name = request.Name
	eventParticipant.Email = request.Email
}
