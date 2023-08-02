package useCase

import (
	"context"
	"github.com/serhiq/effeInventory/services/inventory/internal/domain"
)

type InventoryItem interface {
	Create(c context.Context, items ...*domain.InventoryItem) ([]*domain.InventoryItem, error)
	Update(c context.Context, item domain.InventoryItem) (*domain.InventoryItem, error)
	Delete(c context.Context, ID string) error

	InventoryReader
}

type InventoryReader interface {
	List(c context.Context) ([]*domain.InventoryItem, error)
	ReadByID(c context.Context, ID string) (response *domain.InventoryItem, err error)
}
