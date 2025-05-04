package domain

import "time"

type PaymentStatus string

const (
	PaymentStatusInitial PaymentStatus = "initial"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)

type Payment struct {
	ID            uint          `gorm:"PrimaryKey" json:"id"`
	UserID        uint          `json:"user_id"`
	CaptureMethod string        `json:"capture_method"`
	Amount        float64       `json:"amount"`
	TransactionID uint          `json:"transaction_id"`
	OrderID       string        `json:"order_id"`
	CustomerID    string        `json:"customer_id"` // stripe customer if
	PaymentID     string        `json:"payment_id"`  // payment id
	ClientSecret  string        `json:"client_secret"`
	Status        PaymentStatus `json:"status" gorm:"default:initial"` // initial, success, failed
	Response      string        `json:"response"`
	PaymentUrl    string        `json:"payment_url"`
	CreatedAt     time.Time     `gorm:"default:current_timestamp"`
	UpdatedAt     time.Time     `gorm:"default:current_timestamp"`
}
