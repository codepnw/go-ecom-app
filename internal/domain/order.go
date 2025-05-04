package domain

import "time"

type Order struct {
	ID             uint        `gorm:"PrimaryKey" json:"id"`
	UserID         uint        `json:"user_id"`
	Status         string      `json:"status"`
	Amount         float64     `json:"amount"`
	TransactionID  string      `json:"transaction_id"`
	OrderRefNumber string      `json:"order_ref_number"`
	PaymentID      string      `json:"payment_id"`
	Items          []OrderItem `json:"items"`
	CreatedAt      time.Time   `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time   `gorm:"default:current_timestamp"`
}
