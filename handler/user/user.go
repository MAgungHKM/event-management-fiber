package user

import (
	"database/sql"
	"event-management/handler"
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"

	"github.com/gofiber/fiber/v2"
)

// User godoc
// @Summary     Find All User
// @Description Find all user with their data
// @ID          user-find-all
// @Accept      json
// @Produce     json
// @Tags        User
// @Success     200 {object} UserResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /user [get]
// @Security    ApiBearerToken
func FindAll(c *fiber.Ctx) error {
	users := model.Users{}

	err := users.FindAll()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "Success",
		"data":    users,
	}

	return handler.ApiResponse200(c, resJson)
}

// User godoc
// @Summary     Create a User
// @Description Create a User
// @ID          user-create
// @Accept      json
// @Produce     json
// @Tags        User
// @Param       Body body     UserRequest true "Request Body"
// @Success     201  {object} model.Response
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /user [post]
// @Security    ApiBearerToken
func Create(c *fiber.Ctx) error {
	var request UserRequest
	err := c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	user := model.User{}
	request.MapToUser(&user)

	err = user.Create()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "User created successfully",
	}

	return handler.ApiResponse(c, resJson, fiber.StatusCreated)
}

// User godoc
// @Summary     Update a User
// @Description Update a user with their data
// @ID          user-update
// @Accept      json
// @Produce     json
// @Tags        User
// @Param       id   path     int         true "User ID"
// @Param       Body body     UserRequest true "Request Body"
// @Success     200  {object} UserResponse
// @Failure     400  {object} model.Response
// @Failure     404  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /user/{id} [patch]
// @Security    ApiBearerToken
func Update(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid user ID", fiber.StatusBadRequest))
	}

	user := model.User{ID: *ID}

	err := user.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("User not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	var request UserRequest
	err = c.BodyParser(&request)

	if err != nil {
		panic(err)
	}

	err = utils.ValidateStruct(request)
	if errors.IsValid(err) {
		panic(err)
	}

	request.MapToUser(&user)

	err = user.Update()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "User updated successfully",
	}

	return handler.ApiResponse200(c, resJson)
}

// User godoc
// @Summary     Delete a User
// @Description Delete a user with their data
// @ID          user-delete
// @Accept      json
// @Produce     json
// @Tags        User
// @Param       id  path     int true "User ID"
// @Success     200 {object} UserResponse
// @Failure     400 {object} model.Response
// @Failure     404 {object} model.Response
// @Failure     422 {object} model.ResponseWithError
// @Failure     500 {object} model.Response
// @Router      /user/{id} [delete]
// @Security    ApiBearerToken
func Delete(c *fiber.Ctx) error {
	rawID := c.Params("id")
	ID := utils.ParseInt(rawID)
	if ID == nil {
		panic(errors.New("Invalid user ID", fiber.StatusBadRequest))
	}

	user := model.User{ID: *ID}

	err := user.Find()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("User not found", fiber.StatusNotFound))
		}

		panic(err)
	}

	err = user.Delete()
	if errors.IsValid(err) {
		panic(err)
	}

	resJson := fiber.Map{
		"message": "User deleted successfully",
	}

	return handler.ApiResponse200(c, resJson)
}
