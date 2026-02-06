package operatorUseCase

import (
	"context"
	"fmt"
	"ylanzinhoy-operator-management/internal/domain/entity/operator"
	"ylanzinhoy-operator-management/internal/infrastructure/persistence/storage"

	"golang.org/x/crypto/bcrypt"
)

type Adapter struct {
	Storage storage.Port
}

func (a *Adapter) InsertOperator(ctx context.Context, insertOperatorDto InsertOperatorDTO) (any, error) {

	name, err := operatorEntity.NewName(insertOperatorDto.Name)

	if err != nil {
		return nil, err
	}

	email, err := operatorEntity.NewEmail(insertOperatorDto.Email)

	if err != nil {
		return nil, err
	}

	phone, err := operatorEntity.NewPhone(insertOperatorDto.Phone)
	if err != nil {
		return nil, err
	}

	password, err := operatorEntity.NewPassword(insertOperatorDto.Password)

	if err != nil {
		return nil, err
	}

	passwordHash, err := a.operatorPasswordHash(password.Value)

	if err != nil {
		return nil, err
	}

	address, err := operatorEntity.NewAddress(
		insertOperatorDto.Address.Street,
		insertOperatorDto.Address.Number,
		insertOperatorDto.Address.City,
		insertOperatorDto.Address.State,
		insertOperatorDto.Address.Zip,
	)

	if err != nil {
		return nil, err
	}

	vat, err := operatorEntity.NewVatNumber(insertOperatorDto.VatNumber)
	if err != nil {
		return nil, err
	}

	newOperatorEntity, err := operatorEntity.NewOperator(name, email, phone, passwordHash, address, vat)
	if err != nil {
		return nil, err
	}

	resultId, err := a.Storage.Insert(ctx, newOperatorEntity)

	if err != nil {
		return nil, err
	}

	return resultId, nil

}

func (a *Adapter) ListOperators(ctx context.Context) ([]FetchAllOperatorsOutput, error) {

	operators, err := a.Storage.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing operators: %w", err)
	}

	result := make([]FetchAllOperatorsOutput, 0, len(operators))

	for _, op := range operators {
		output := FetchAllOperatorsOutput{
			Name:      op.Name.Value,
			Email:     op.Email.Value,
			Phone:     op.Phone.Value,
			VatNumber: op.VatNumber.Value,
			Address: AddressInput{
				Street: op.Address.Street,
				Number: op.Address.Number,
				City:   op.Address.City,
				State:  op.Address.State,
				Zip:    op.Address.Zip,
			},
		}

		result = append(result, output)
	}

	return result, nil
}

func (a *Adapter) GetOperatorById(ctx context.Context, id string) (*operatorEntity.Operator, error) {
	operator, err := a.Storage.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting operator by id: %w", err)
	}
	return operator, nil
}

func (a *Adapter) UpdateOperator(ctx context.Context, id string, updateOperatorDto UpdateOperatorDTO) error {
	operator, err := a.Storage.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("operator not found: %w", err)
	}

	name, err := operatorEntity.NewName(updateOperatorDto.Name)
	if err != nil {
		return err
	}
	operator.Name = name

	email, err := operatorEntity.NewEmail(updateOperatorDto.Email)
	if err != nil {
		return err
	}
	operator.Email = email

	phone, err := operatorEntity.NewPhone(updateOperatorDto.Phone)
	if err != nil {
		return err
	}
	operator.Phone = phone

	address, err := operatorEntity.NewAddress(
		updateOperatorDto.Address.Street,
		updateOperatorDto.Address.Number,
		updateOperatorDto.Address.City,
		updateOperatorDto.Address.State,
		updateOperatorDto.Address.Zip,
	)
	if err != nil {
		return err
	}
	operator.Address = address

	vat, err := operatorEntity.NewVatNumber(updateOperatorDto.VatNumber)
	if err != nil {
		return err
	}
	operator.VatNumber = vat

	err = a.Storage.UpdateByID(ctx, id, operator)
	if err != nil {
		return fmt.Errorf("error updating operator: %w", err)
	}

	return nil
}

func (a *Adapter) DeleteOperator(ctx context.Context, id string) error {
	err := a.Storage.DeleteByID(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting operator: %w", err)
	}
	return nil
}

func (a *Adapter) operatorPasswordHash(passwordString string) (operatorEntity.Password, error) {
	newPass, err := operatorEntity.NewPassword(passwordString)

	if err != nil {
		return operatorEntity.Password{}, fmt.Errorf("error creating password: %w", err)

	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPass.Value), 14)

	if err != nil {
		return operatorEntity.Password{}, fmt.Errorf("error hashing password: %w", err)
	}

	newPass.Value = string(hash)

	return newPass, nil
}

func (a *Adapter) UpdatePassword(ctx context.Context, input UpdatePasswordDTO) error {

	OperatorFinded, err := a.Storage.FindByID(ctx, input.ID)

	if err != nil {
		return err
	}

	newPasswordHash, err := a.operatorPasswordHash(input.NewPassword)

	if err != nil || newPasswordHash.Value == "" {
		return err
	}

	if err := OperatorFinded.ChangePassword(newPasswordHash); err != nil {
		return err
	}

	return a.Storage.UpdatePassword(ctx, input.ID, OperatorFinded)
}
