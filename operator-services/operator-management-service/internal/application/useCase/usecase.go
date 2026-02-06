package useCase

import (
	"ylanzinhoy-operator-management/internal/application/useCase/operatorUseCase"
	"ylanzinhoy-operator-management/internal/infrastructure/persistence/storage"
)

func NewOperatorUseCase(storage storage.Port) *operatorUseCase.Adapter {
	return &operatorUseCase.Adapter{
		Storage: storage,
	}
}
