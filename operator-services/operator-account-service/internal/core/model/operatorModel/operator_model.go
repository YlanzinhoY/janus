package operatormodel

import (
	"time"
)

// mongodb
type OperatorModel struct {
	uid       string    `bson:"uid"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Phone     string    `bson:"phone"`
	Password  string    `bson:"password"`
	Address   Address   `bson:"address"`
	VatNumber string    `bson:"vatNumber"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func NewOperatorModel(
	uid string,
	name string,
	email string,
	phone string,
	address Address,
	vatNumber string,
	createdAt time.Time,
	updatedAt time.Time) *OperatorModel {
	return &OperatorModel{
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		VatNumber: vatNumber,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt}
}

type Address struct {
	Street string `bson:"street"`
	City   string `bson:"city"`
	State  string `bson:"state"`
	Number string `bson:"numero"`
	Zip    string `bson:"zip"`
}
