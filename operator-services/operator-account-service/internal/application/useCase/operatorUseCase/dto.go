package operatorUseCase

type InsertOperatorDTO struct {
	uid       string       `validate:"required" json:"uid"`
	Name      string       `validate:"required" json:"name"`
	Email     string       `validate:"required,email" json:"email"`
	Phone     string       `validate:"required" json:"phone"`
	Password  string       `validate:"required,validPassword" json:"password"`
	Address   AddressInput `validate:"required" json:"address"`
	VatNumber string       `validate:"required,validVatNumber" json:"vatNumber"`
}

type AddressInput struct {
	Street string `validate:"required" json:"street"`
	Number string `validate:"required" json:"number"`
	City   string `validate:"required" json:"city"`
	State  string `validate:"required" json:"state"`
	Zip    string `validate:"required" json:"zip"`
}

type FetchAllOperatorsOutput struct {
	uid       string       `json:"uid"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Phone     string       `json:"phone"`
	Address   AddressInput `json:"address"`
	VatNumber string       `json:"vatNumber"`
}

type UpdateOperatorDTO struct {
	Name      string        `json:"name,omitempty"`
	Email     string        `validate:"omitempty,email" json:"email,omitempty"`
	Phone     string        `json:"phone,omitempty"`
	Address   *AddressInput `json:"address,omitempty"`
	VatNumber string        `validate:"omitempty,validVatNumber" json:"vatNumber,omitempty"`
}

type UpdatePasswordDTO struct {
	ID          string `validate:"required" json:"id"`
	NewPassword string `validate:"required,validPassword" json:"newPassword"`
}
