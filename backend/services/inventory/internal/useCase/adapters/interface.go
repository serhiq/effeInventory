package storage

import (
	"context"
	"github.com/serhiq/effeInventory/services/inventory/internal/domain"
)

type Storage interface {
	Inventory
}

type Inventory interface {
	CreateInventory(ctx context.Context, items ...*domain.InventoryItem) ([]*domain.InventoryItem, error)
	UpdateInventory(ctx context.Context, ID string, updateFn func(c *domain.InventoryItem) (*domain.InventoryItem, error)) (*domain.InventoryItem, error)
	DeleteInventory(ctx context.Context, ID string) error

	InventoryReader
}

type InventoryReader interface {
	ListInventoryItem(ctx context.Context) ([]*domain.InventoryItem, error)
	// todo adding filter
	//ListInventoryItem(ctx context.Context, parameter queryParameter.QueryParameter) ([]*domain.InventoryItem, error)
	ReadInventoryByID(ctx context.Context, ID string) (response *domain.InventoryItem, err error)
}
