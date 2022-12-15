package errors

import (
	"errors"
	"event-management/handler"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var Default = fiber.Error{Message: "Server Error", Code: 500}

var ErrorHandler = func(c *fiber.Ctx, err error) error {
	// Default 500 statuscode
	code := fiber.StatusInternalServerError
	data := fiber.Map{"message": err.Error()}

	if e, ok := err.(*fiber.Error); ok {
		// Override status code & data if fiber.Error type
		code = e.Code
		data = fiber.Map{"message": e.Message}
	}

	if e, ok := err.(*ErrorValidations); ok {
		// Override status code & data if ErrorValidations type
		code = fiber.StatusUnprocessableEntity
		data = fiber.Map{
			"message": "Request Body Error",
			"error":   e,
		}
	}

	if e := err.Error(); strings.Contains(e, "JWT") || strings.Contains(e, "Token") || strings.Contains(e, "signature") {
		// Override status code & data if JWT type error
		message := e

		if strings.Contains(e, "JWT") || strings.Contains(e, "signature") {
			message = "Unauthorized"
		}

		code = fiber.StatusUnauthorized
		data = fiber.Map{"message": message}
	}

	return handler.ApiResponse(c, data, code)
}

func New(message string, code int) *fiber.Error {
	return &fiber.Error{
		Message: message,
		Code:    code,
	}
}

func Error422(c *fiber.Ctx, err error) error {
	panic(New(err.Error(), fiber.StatusUnprocessableEntity))
}

func Error404(c *fiber.Ctx) error {
	panic(New("Not Found", fiber.StatusNotFound))
}

func IsValid(err error) bool {
	return err != nil && err.Error() != ""
}

// Is reports whether any error in err's chain matches target.
// The chain consists of err itself followed by the sequence of errors obtained by
// repeatedly calling Unwrap.
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
// An error type might provide an Is method so it can be treated as equivalent
// to an existing error. For example, if MyError defines
//
//	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
//
// then Is(MyError{}, fs.ErrExist) returns true. See syscall.Errno.Is for
// an example in the standard library. An Is method should only shallowly
// compare err and the target and not call Unwrap on either.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}
