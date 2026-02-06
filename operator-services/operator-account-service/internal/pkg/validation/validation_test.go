package validation_test

import (
	"testing"

	"ylanzinhoy-operator-management/internal/pkg/validation"

	"github.com/stretchr/testify/assert"
)

type UpdatePasswordDTO struct {
	ID          string `validate:"required" json:"id"`
	NewPassword string `validate:"required,validPassword" json:"newPassword"`
}

func TestPasswordValidation(t *testing.T) {
	t.Run("valid passwords", func(t *testing.T) {
		validPasswords := []struct {
			name     string
			password string
		}{
			{"with special chars and numbers", "ValidPass@123"},
			{"with underscore", "MyP@ssw0rd_Test"},
			{"with hash", "Test#123Abc"},
			{"with exclamation", "Secure!Pass99"},
			{"minimum length with all requirements", "aA1!bcde"},
		}

		for _, tc := range validPasswords {
			t.Run(tc.name, func(t *testing.T) {
				dto := UpdatePasswordDTO{
					ID:          "123",
					NewPassword: tc.password,
				}

				_, errMsg := validation.ValidateStruct(dto)
				assert.Empty(t, errMsg, "password %q should be valid", tc.password)
			})
		}
	})

	t.Run("invalid passwords - missing uppercase", func(t *testing.T) {
		invalidPasswords := []string{
			"validpass@123",
			"lowercase#99",
			"test!password1",
		}

		for _, password := range invalidPasswords {
			t.Run(password, func(t *testing.T) {
				dto := UpdatePasswordDTO{
					ID:          "123",
					NewPassword: password,
				}

				_, errMsg := validation.ValidateStruct(dto)
				assert.NotEmpty(t, errMsg, "should fail without uppercase")
				assert.Contains(t, errMsg, "NewPassword")
				assert.Contains(t, errMsg, "validPassword")
			})
		}
	})

	t.Run("invalid passwords - missing lowercase", func(t *testing.T) {
		invalidPasswords := []string{
			"VALIDPASS@123",
			"UPPERCASE#99",
			"TEST!PASSWORD1",
		}

		for _, password := range invalidPasswords {
			t.Run(password, func(t *testing.T) {
				dto := UpdatePasswordDTO{
					ID:          "123",
					NewPassword: password,
				}

				_, errMsg := validation.ValidateStruct(dto)
				assert.NotEmpty(t, errMsg, "should fail without lowercase")
				assert.Contains(t, errMsg, "NewPassword")
				assert.Contains(t, errMsg, "validPassword")
			})
		}
	})

	t.Run("invalid passwords - missing special character", func(t *testing.T) {
		invalidPasswords := []string{
			"ValidPass123",
			"MyPassword99",
			"TestAbc123",
		}

		for _, password := range invalidPasswords {
			t.Run(password, func(t *testing.T) {
				dto := UpdatePasswordDTO{
					ID:          "123",
					NewPassword: password,
				}

				_, errMsg := validation.ValidateStruct(dto)
				assert.NotEmpty(t, errMsg, "should fail without special character")
				assert.Contains(t, errMsg, "NewPassword")
				assert.Contains(t, errMsg, "validPassword")
			})
		}
	})

	t.Run("invalid passwords - too short", func(t *testing.T) {
		invalidPasswords := []string{
			"aA1!",
			"Sh0rt!",
			"Te$t1",
		}

		for _, password := range invalidPasswords {
			t.Run(password, func(t *testing.T) {
				dto := UpdatePasswordDTO{
					ID:          "123",
					NewPassword: password,
				}

				_, errMsg := validation.ValidateStruct(dto)
				assert.NotEmpty(t, errMsg, "should fail - too short")
				assert.Contains(t, errMsg, "NewPassword")
				assert.Contains(t, errMsg, "validPassword")
			})
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		t.Run("missing ID", func(t *testing.T) {
			dto := UpdatePasswordDTO{
				ID:          "",
				NewPassword: "ValidPass@123",
			}

			_, errMsg := validation.ValidateStruct(dto)
			assert.NotEmpty(t, errMsg)
			assert.Contains(t, errMsg, "ID")
			assert.Contains(t, errMsg, "required")
		})

		t.Run("missing password", func(t *testing.T) {
			dto := UpdatePasswordDTO{
				ID:          "123",
				NewPassword: "",
			}

			_, errMsg := validation.ValidateStruct(dto)
			assert.NotEmpty(t, errMsg)
			assert.Contains(t, errMsg, "NewPassword")
			assert.Contains(t, errMsg, "required")
		})

		t.Run("missing both", func(t *testing.T) {
			dto := UpdatePasswordDTO{
				ID:          "",
				NewPassword: "",
			}

			_, errMsg := validation.ValidateStruct(dto)
			assert.NotEmpty(t, errMsg)
			// Vai falhar no primeiro campo (ID)
			assert.Contains(t, errMsg, "ID")
		})
	})
}
