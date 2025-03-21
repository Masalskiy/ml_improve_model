package models

type Product struct {
	Name        string `json:"name"`
	SKU         string `json:"sku"`
	Link        string `json:"link"`
	ImageLink   string `json:"image_link"`
	Description string `json:"description"`
	IDProduct   int64  `json:"id_product"`
} 