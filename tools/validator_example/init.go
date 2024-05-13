package validator_example

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate = validator.New()

func init() {
	registerValidation(isValidAgeTag, validAge)
	registerValidation(isValidHobbyTag, validHobby)
}

func registerValidation(tag string, fn validator.Func) {
	if err := validate.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("register validator for '%s' error: %v", tag, err)
	}
}

const (
	isValidAgeTag   = "isValidAge"
	isValidHobbyTag = "isValidHobby"
)

var (
	InvalidAgeErr   = errors.New("age should be between 0 and 200")
	InvalidHobbyErr = errors.New("hobby should be 'ball' or 'swim'")
)

func validateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	// refer to: https://github.com/go-playground/validator/issues/559
	for _, validateErr := range err.(validator.ValidationErrors) {
		// only check first error
		switch validateErr.Tag() {
		case isValidAgeTag:
			// return custom error message
			return InvalidAgeErr
		case isValidHobbyTag:
			return InvalidHobbyErr
		default:
			return err
		}
	}

	return err
}

func validAge(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	if age > 0 && age < 200 {
		return true
	}
	return false
}

func validHobby(fl validator.FieldLevel) bool {
	hobby := fl.Field().String()
	if hobby == "ball" || hobby == "swim" {
		return true
	}
	return false
}
