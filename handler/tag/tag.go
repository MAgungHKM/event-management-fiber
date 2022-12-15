package tag

import (
	"database/sql"
	"event-management/handler"
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"

	"github.com/gofiber/fiber/v2"
)

// Tag godoc
// @Summary     Find All Tag
// @Description Find all Tag with their data
// @ID          tag-find-all
// @Accept      json
// @Produce     json
// @Tags        Tag
// @Success     200 {object} TagResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /tag [get]
// @Security    ApiBearerToken
func FindAll(c *fiber.Ctx) error {
	tags := model.Tags{}

	err := tags.FindAll()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Success",
		"data":    tags,
	}

	return handler.ApiResponse200(c, resJson)
}

// Tag godoc
// @Summary     Create a Tag
// @Description Create a Tag
// @ID          tag-create
// @Accept      json
// @Produce     json
// @Tags        Tag
// @Param       Body body     TagRequest true "Request Body"
// @Success     201  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /tag [post]
// @Security    ApiBearerToken
func Create(c *fiber.Ctx) error {
	var request TagRequest
	err := c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	tag := model.Tag{}
	request.MapToTag(&tag)

	err = tag.Create()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Tag created successfully",
	}

	return handler.ApiResponse(c, resJson, fiber.StatusCreated)
}

// Tag godoc
// @Summary     Update a Tag
// @Description Update a Tag with their data
// @ID          tag-update
// @Accept      json
// @Produce     json
// @Tags        Tag
// @Param       id   path     int        true "Tag ID"
// @Param       Body body     TagRequest true "Request Body"
// @Success     200  {object} TagResponse
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /tag/{id} [patch]
// @Security    ApiBearerToken
func Update(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid tag ID", fiber.StatusBadRequest))
	}

	tag := model.Tag{ID: *ID}

	err := tag.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Tag not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	var request TagRequest
	err = c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	request.MapToTag(&tag)

	err = tag.Update()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Tag updated successfully",
	}

	return handler.ApiResponse200(c, resJson)
}

// Tag godoc
// @Summary     Delete a Tag
// @Description Delete a Tag with their data
// @ID          tag-delete
// @Accept      json
// @Produce     json
// @Tags        Tag
// @Param       id  path     int true "Tag ID"
// @Success     200 {object} TagResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /tag/{id} [delete]
// @Security    ApiBearerToken
func Delete(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid tag ID", fiber.StatusBadRequest))
	}

	tag := model.Tag{ID: *ID}

	err := tag.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Tag not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	err = tag.Delete()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Tag deleted successfully",
	}

	return handler.ApiResponse200(c, resJson)
}
