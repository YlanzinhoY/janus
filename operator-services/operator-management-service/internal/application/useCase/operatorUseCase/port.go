package operatorUseCase

import (
	"context"
	operatorEntity "ylanzinhoy-operator-management/internal/domain/entity/operator"
)

type Port interface {
	InsertOperator(ctx context.Context, insertOperatorDto InsertOperatorDTO) (any, error)
	ListOperators(ctx context.Context) ([]FetchAllOperatorsOutput, error)
	GetOperatorById(ctx context.Context, id string) (*operatorEntity.Operator, error)
	UpdateOperator(ctx context.Context, id string, updateOperatorDto UpdateOperatorDTO) error
	DeleteOperator(ctx context.Context, id string) error

	operatorPasswordHash(passwordString string) (operatorEntity.Password, error)

	UpdatePassword(ctx context.Context, input UpdatePasswordDTO) error
}
