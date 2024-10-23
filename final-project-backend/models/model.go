package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Password     string         `json:"password"`
	Role         string         `json:"role" gorm:"type:enum('admin', 'user') default:'user'; not null"`
	Balance      float64        `json:"balance"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Transactions []Transactions
}

type RecyclableItems struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	PricePerKg  float64        `json:"price_per_kg"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Transactions struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Quantity        int       `json:"quantity"`
	TotalPrice      float64   `json:"total_price"`
	OrderDate       time.Time `json:"order_at"`
	UserID          uint      `json:"user_id"`
	RecycleItemID   uint      `json:"recycle_item_id"`
	RecycleAddresID uint      `json:"recycle_addres_id"`
	RecyclableItems []RecyclableItems
	RecyclingCenter RecyclingCenter
}

type RecyclingCenter struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Addres        string         `json:"addres"`
	ContactNumber string         `json:"contact_number"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	TransactionId uint           `json:"transaction_id"`
}
