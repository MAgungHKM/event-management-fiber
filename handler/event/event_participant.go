package event

import (
	"database/sql"
	"event-management/handler"
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

// Event godoc
// @Summary     Create a participant for an Event
// @Description Create a participant with their data for an Event
// @ID          event-create-participant
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       id   path     int    true "Event ID"
// @Param       Body body     EventParticipantRequest true "Request Body"
// @Success     200  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     403  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /event/{id}/participate [post]
func Participate(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid event ID", fiber.StatusBadRequest))
	}

	event := model.Event{ID: *ID}

	err := event.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Event not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	var request EventParticipantRequest
	err = c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	eventParticipant := model.EventParticipant{EventID: event.ID}
	request.MapToEvent(&eventParticipant)

	err = eventParticipant.FindByEmailAndEventID()
	if errors.IsValid(err) && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}

	if eventParticipant.Status == "ATTENDED" {
		panic(errors.New("You have attended this event", fiber.StatusForbidden))
	}

	if eventParticipant.Status == "REGISTERED" {
		panic(errors.New("You have registered to this event", fiber.StatusForbidden))
	}

	err = eventParticipant.Create()
	if errors.IsValid(err) {
		panic(err)
	}

	var png []byte
	png, err = qrcode.Encode(fmt.Sprintf("%d@%d-", event.ID, eventParticipant.ID), qrcode.Medium, 256)
	if errors.IsValid(err) {
		panic(err)
	}

	c.Set(fiber.HeaderContentType, "image/png")
	c.Status(fiber.StatusOK)
	return c.Send(png)
}

// Event godoc
// @Summary     Attend an Event
// @Description Attend an Event
// @ID          event-attend
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       id   path     int    true "Event ID"
// @Param       code path     string true "Participant Code"
// @Success     200  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     403  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /event/{id}/attend/{code} [patch]
// @Security    ApiBearerToken
func Attend(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid event ID", fiber.StatusBadRequest))
	}

	event := model.Event{ID: *ID}

	err := event.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Event not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	rawCode := c.Params("code")
	code, err := url.QueryUnescape(rawCode)
	if err != nil {
		panic(errors.New("Invalid participant code", fiber.StatusBadRequest))
	}

	codeArr := strings.Split(code, "@")
	rawParticipantID := codeArr[1][:len(codeArr[1])-1]
	participantID := utils.ParseInt(rawParticipantID)

	if participantID == nil {
		panic(errors.New("Invalid participant code", fiber.StatusBadRequest))
	}

	eventParticipant := model.EventParticipant{ID: *participantID}

	err = eventParticipant.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Participant not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	if eventParticipant.Status == "ATTENDED" {
		panic(errors.New("You have attended this event", fiber.StatusForbidden))
	}

	if eventParticipant.Status == "CANCELED" {
		panic(errors.New("You have canceled your participation to this event", fiber.StatusForbidden))
	}

	err = eventParticipant.Attend()
	if errors.IsValid(err) {
		panic(err)
	}

	return handler.ApiResponse200(c, fiber.Map{"message": "The event has been successfully attended"})
}

// Event godoc
// @Summary     Cancel participation of an Event
// @Description Cancel participation of an Event
// @ID          event-cancel
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       id   path     int                     true "Event ID"
// @Param       code path     string true "Participant Code"
// @Success     200  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     403  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /event/{id}/cancel/{code} [delete]
func Cancel(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid event ID", fiber.StatusBadRequest))
	}

	event := model.Event{ID: *ID}

	err := event.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Event not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	rawCode := c.Params("code")
	code, err := url.QueryUnescape(rawCode)
	if err != nil {
		panic(errors.New("Invalid participant code", fiber.StatusBadRequest))
	}

	codeArr := strings.Split(code, "@")
	rawParticipantID := codeArr[1][:len(codeArr[1])-1]
	participantID := utils.ParseInt(rawParticipantID)

	if participantID == nil {
		panic(errors.New("Invalid participant code", fiber.StatusBadRequest))
	}

	eventParticipant := model.EventParticipant{ID: *participantID}

	err = eventParticipant.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Participant not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	if eventParticipant.Status == "ATTENDED" {
		panic(errors.New("You have attended this event", fiber.StatusForbidden))
	}

	err = eventParticipant.Cancel()
	if errors.IsValid(err) {
		panic(err)
	}

	return handler.ApiResponse200(c, fiber.Map{"message": "The event participation has been successfully canceled"})
}
