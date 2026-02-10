package operatormodel

import (
	"fmt"
	operatorEntity "ylanzinhoy-operator-management/internal/domain/entity/operator"

	"github.com/google/uuid"
)

func ModelToEntity(model OperatorModel) (*operatorEntity.Operator, error) {

	uid := uuid.New()

	email, err := operatorEntity.NewEmail(model.Email)
	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (email): %w", err)
	}

	name, err := operatorEntity.NewName(model.Name)
	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (name): %w", err)
	}

	phone, err := operatorEntity.NewPhone(model.Phone)
	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (phone): %w", err)
	}

	address, err := operatorEntity.NewAddress(
		model.Address.Street,
		model.Address.Number,
		model.Address.City,
		model.Address.State,
		model.Address.Zip,
	)

	password, err := operatorEntity.NewPassword(model.Password)

	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (password): %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (address): %w", err)
	}

	vatNumber, err := operatorEntity.NewVatNumber(model.VatNumber)
	if err != nil {
		return nil, fmt.Errorf("fail to convert model to entity (vatNumber): %w", err)
	}

	operator, err := operatorEntity.NewOperator(uid, name, email, phone, password, address, vatNumber)
	if err != nil {
		return nil, fmt.Errorf("fail to create operator entity: %w", err)
	}

	return operator, nil
}
func EntityToModel(entity *operatorEntity.Operator) *OperatorModel {
	return &OperatorModel{
		uid:      entity.Uid.String(),
		Name:     entity.Name.Value,
		Email:    entity.Email.Value,
		Phone:    entity.Phone.Value,
		Password: entity.Password.Value,
		Address: Address{
			Street: entity.Address.Street,
			City:   entity.Address.City,
			State:  entity.Address.State,
			Number: entity.Address.Number,
			Zip:    entity.Address.Zip,
		},
		VatNumber: entity.VatNumber.Value,
	}
}
