package operatorEntity

import (
	"fmt"
	"ylanzinhoy-operator-management/internal/pkg/validation"
)

type VatNumber struct {
	Value string `validate:"required,validVatNumber"`
}

func NewVatNumber(value string) (VatNumber, error) {
	v := VatNumber{Value: value}
	if err := validation.GetInstance().Struct(v); err != nil {
		return VatNumber{}, fmt.Errorf("invalid vat number: %w", err)
	}
	return v, nil
}

type Email struct {
	Value string `validate:"required,email"`
}

func NewEmail(value string) (Email, error) {
	e := Email{Value: value}
	if err := validation.GetInstance().Struct(e); err != nil {
		return Email{}, fmt.Errorf("invalid email: %w", err)
	}
	return e, nil
}

type Phone struct {
	Value string `validate:"required"`
}

func NewPhone(value string) (Phone, error) {
	p := Phone{Value: value}
	if err := validation.GetInstance().Struct(p); err != nil {
		return Phone{}, fmt.Errorf("invalid phone: %w", err)
	}
	return p, nil
}

type Address struct {
	Street string `validate:"required"`
	Number string `validate:"required"`
	City   string `validate:"required"`
	State  string `validate:"required"`
	Zip    string `validate:"required"`
}

func NewAddress(street, number, city, state, zip string) (Address, error) {
	a := Address{
		Street: street,
		Number: number,
		City:   city,
		State:  state,
		Zip:    zip,
	}
	validStruct, errorMsg := validation.ValidateStruct(a)

	if errorMsg != "" {
		return Address{}, fmt.Errorf("invalid address: %s", errorMsg)
	}

	return validStruct, nil
}

type Name struct {
	Value string `validate:"required"`
}

func NewName(value string) (Name, error) {
	n := Name{Value: value}
	if err := validation.GetInstance().Struct(n); err != nil {
		return Name{}, fmt.Errorf("invalid name: %w", err)
	}
	return n, nil
}

type Password struct {
	Value string `validate:"required,validPassword"`
}

func NewPassword(value string) (Password, error) {
	pass := Password{Value: value}

	passValidation, errMsg := validation.ValidateStruct(pass)

	if errMsg != "" {
		return Password{}, fmt.Errorf("invalid password: %s", errMsg)
	}

	return passValidation, nil
}
