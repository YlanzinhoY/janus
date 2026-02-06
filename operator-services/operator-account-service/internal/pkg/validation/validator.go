package validation

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetInstance() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		if err := validate.RegisterValidation("validVatNumber", VatNumberValidation); err != nil {
			panic(fmt.Sprintf("Failed to register validVatNumber: %v", err))
		}

		if err := validate.RegisterValidation("validPassword", PasswordValidation); err != nil {
			panic(fmt.Sprintf("Failed to register validPassword: %v", err))
		}
	})

	return validate
}

func ValidateStruct[T any](s T) (T, string) {
	v := GetInstance()

	if err := v.Struct(s); err != nil {
		return s, formatError(err)
	}

	return s, ""
}

func formatError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return fmt.Sprintf("field %s error on validation: %s", validationErrors[0].Field(), validationErrors[0].Tag())
	}

	return err.Error()
}
