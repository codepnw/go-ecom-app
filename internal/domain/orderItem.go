package domain

import "time"

type OrderItem struct {
	ID        uint      `gorm:"PrimaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	SellerID  uint      `json:"seller_id"`
	Price     float64    `json:"price"`
	Qty       uint      `json:"qty"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
