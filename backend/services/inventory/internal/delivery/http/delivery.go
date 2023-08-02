package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/serhiq/effeInventory/services/inventory/internal/useCase"
)

type Delivery struct {
	ucInventory useCase.InventoryItem
	router      *gin.Engine

	options Options
}

type Options struct {
	Port uint
}

func New(ucInventory useCase.InventoryItem, options Options) *Delivery {
	var d = &Delivery{
		ucInventory: ucInventory,
	}

	d.SetOptions(options)

	d.router = d.initRouter()
	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(fmt.Sprintf(":%d", uint16(d.options.Port)))
}
