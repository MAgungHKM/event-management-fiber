package event

import (
	"database/sql"
	"event-management/handler"
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Event godoc
// @Summary     Find All Event
// @Description Find all Event with its data
// @ID          event-find-all
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       tag query    string false "Filter with Tag"
// @Success     200 {object} EventResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /event [get]
func FindAll(c *fiber.Ctx) error {
	tag := c.Query("tag")

	events := model.Events{}

	if len(tag) > 0 {
		err := events.FindAllWithTag(tag)
		if errors.IsValid(err) {
			panic(err)
		}
	} else {
		err := events.FindAll()
		if errors.IsValid(err) {
			panic(err)
		}
	}

	resJson := fiber.Map{
		"message": "Success",
		"data":    events,
	}

	return handler.ApiResponse200(c, resJson)
}

// Event godoc
// @Summary     Create an Event
// @Description Create an Event
// @ID          event-create
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       Body body     EventRequest true "Request Body"
// @Success     201  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /event [post]
// @Security    ApiBearerToken
func Create(c *fiber.Ctx) error {
	var request EventRequest
	err := c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	event := model.Event{}
	err = request.MapToEvent(&event)
	if err != nil {
		panic(err)
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	user := model.User{Username: username}
	err = user.FindByUsername()
	if errors.IsValid(err) {
		panic(err)
	}

	event.CreatedBy = user.ID
	event.ContactPerson = user

	err = event.Create()
	if errors.IsValid(err) {
		panic(err)
	}

	for _, tagName := range request.Tags {
		tag := model.Tag{Name: tagName}
		err := tag.FindByName()
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				panic(err)
			}

			err = tag.Create()
			if err != nil {
				panic(err)
			}
		}

		eventTag := model.EventTag{
			EventID: event.ID,
			TagID:   tag.ID,
		}

		err = eventTag.Create()
		if err != nil {
			panic(err)
		}
	}

	event.Tags = request.Tags

	resJson := fiber.Map{
		"message": "Event created successfully",
	}

	return handler.ApiResponse(c, resJson, fiber.StatusCreated)
}

// Event godoc
// @Summary     Update an Event
// @Description Update an Event with their data
// @ID          event-update
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       id   path     int          true "Event ID"
// @Param       Body body     EventRequest true "Request Body"
// @Success     200  {object} EventResponse
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /event/{id} [patch]
// @Security    ApiBearerToken
func Update(c *fiber.Ctx) error {
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

	var request EventRequest
	err = c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	err = request.MapToEvent(&event)
	if err != nil {
		panic(err)
	}

	err = event.Update()
	if errors.IsValid(err) {
		panic(err)
	}

	eventTag := model.EventTag{EventID: event.ID}
	err = eventTag.DeleteByEventID()
	if err != nil {
		panic(err)
	}

	for _, tagName := range request.Tags {
		tag := model.Tag{Name: tagName}
		err := tag.FindByName()
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				panic(err)
			}

			err = tag.Create()
			if err != nil {
				panic(err)
			}
		}

		eventTag := model.EventTag{
			EventID: event.ID,
			TagID:   tag.ID,
		}

		err = eventTag.Create()
		if err != nil {
			panic(err)
		}
	}

	event.Tags = request.Tags

	resJson := fiber.Map{
		"message": "Event updated successfully",
	}

	return handler.ApiResponse200(c, resJson)
}

// Event godoc
// @Summary     Delete an Event
// @Description Delete an Event with their data
// @ID          event-delete
// @Accept      json
// @Produce     json
// @Tags        Event
// @Param       id  path     int true "Event ID"
// @Success     200 {object} EventResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /event/{id} [delete]
// @Security    ApiBearerToken
func Delete(c *fiber.Ctx) error {
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

	err = event.Delete()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Event deleted successfully",
	}

	return handler.ApiResponse200(c, resJson)
}
