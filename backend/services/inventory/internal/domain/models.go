package domain

type InventoryItem struct {
	Uuid       string   `json:"uuid"`
	Name       string   `json:"name"`
	Barcodes   []string `json:"barcodes"`
	Quantity   int      `json:"quantity"`
	Tags       []string `json:"tags"`
	ChildUuids []string `json:"childUuids"`
	ImageUrls  []string `json:"imageUrls"`
	DocUrls    []string `json:"docUrls"`
}
