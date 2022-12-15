package auth

import (
	"database/sql"
	"event-management/handler"
	"event-management/handler/errors"
	"event-management/model"
	"event-management/utils"
	"event-management/utils/secret"

	"github.com/gofiber/fiber/v2"
)

// Auth godoc
// @Summary     Authenticate for JWT
// @Description Generate JWT if authentication is successful
// @ID          auth
// @Accept      json
// @Produce     json
// @Tags        Auth
// @Param       Body body     AuthRequest true "Request Body"
// @Success     200  {object} AuthResponse
// @Failure     400  {object} model.Response
// @Failure     403  {object} model.Response
// @Failure     422  {object} model.ResponseWithError
// @Failure     500  {object} model.Response
// @Router      /auth [post]
func Auth(c *fiber.Ctx) error {
	var request AuthRequest
	err := c.BodyParser(&request)

	if err == nil {
		err = utils.ValidateStruct(request)
	}

	if errors.IsValid(err) {
		panic(err)
	}

	user := model.User{Username: request.Username}
	err = user.FindByUsername()
	if errors.IsValid(err) {
		if errors.Is(err, sql.ErrNoRows) {
			panic(errors.New("Invalid username or password", fiber.StatusForbidden))
		}

		panic(err)
	}

	isCorrectSecret := secret.VerifyHash(request.Password, user.Secret)

	if !isCorrectSecret {
		panic(errors.New("Invalid password", fiber.StatusBadRequest))
	}

	signedToken, err := request.GenerateAccessToken()
	if errors.IsValid(err) {
		panic(err)
	}

	res := AuthResponse{
		TokenType:   "Bearer",
		AccessToken: signedToken,
	}

	return handler.ApiResponse200(c, res)
}
