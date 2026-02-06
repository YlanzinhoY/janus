package storage

import (
	"context"
	operatorEntity "ylanzinhoy-operator-management/internal/domain/entity/operator"
)

type Port interface {
	Insert(ctx context.Context, entity *operatorEntity.Operator) (any, error)
	FindByID(ctx context.Context, id string) (*operatorEntity.Operator, error)
	FindAll(ctx context.Context) ([]*operatorEntity.Operator, error)
	UpdateByID(ctx context.Context, id string, entity *operatorEntity.Operator) error
	DeleteByID(ctx context.Context, id string) error

	UpdatePassword(ctx context.Context, id string, entity *operatorEntity.Operator) error
}
