package domain

import "time"

type Address struct {
	ID            uint      `json:"id"`
	AddressInput1 string    `json:"address1"`
	AddressInput2 string    `json:"address2"`
	City          string    `json:"city"`
	PostCode      uint      `json:"post_code"`
	Country       string    `json:"country"`
	UserID        uint      `json:"user_id"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
