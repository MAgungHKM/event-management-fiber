package utils

import (
	"event-management/handler/errors"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func IsISO8601DateTime(fl validator.FieldLevel) bool {
	ISO8601DateTimeRegexString := "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
	ISO8601DateTimeRegex := regexp.MustCompile(ISO8601DateTimeRegexString)
	return ISO8601DateTimeRegex.MatchString(fl.Field().String())
}

func ValidateStruct(class interface{}) *errors.ErrorValidations {
	var errorList errors.ErrorValidations
	validate := validator.New()

	err := validate.RegisterValidation("ISO8601datetime", IsISO8601DateTime)
	if err != nil {
		panic(err)
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	err = validate.Struct(class)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := errors.ErrorValidation{
				Field:  err.Field(),
				Reason: err.Tag(),
				Value:  err.Param(),
			}
			errorList = append(errorList, &element)
		}
	}

	return &errorList
}
