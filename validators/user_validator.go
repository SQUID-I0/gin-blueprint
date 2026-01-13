package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ValidateUsername validator.Func = func(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	return match
}

var ValidateStrongPassword validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && len(password) >= 8
}
