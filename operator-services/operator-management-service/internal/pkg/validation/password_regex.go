package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	passHasUpper   = regexp.MustCompile(`[A-Z]`)
	passHasLower   = regexp.MustCompile(`[a-z]`)
	passHasSpecial = regexp.MustCompile(`[^A-Za-z0-9]`) // Isso INCLUI underscore
)

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Mínimo 8 caracteres
	if len(password) < 8 {
		return false
	}

	// Verifica cada requisito
	return passHasUpper.MatchString(password) &&
		passHasLower.MatchString(password) &&
		passHasSpecial.MatchString(password)
}
