package http

import (
	storage "github.com/serhiq/effeInventory/services/inventory/internal/useCase/adapters"
)

type UseCase struct {
	adapterStorage storage.Inventory
}
