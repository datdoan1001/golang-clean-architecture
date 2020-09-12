package model

import "time"

// User struct
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Age       string    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// CreditCard struct
type CreditCard struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	UserID uint   `json:"userid"`
	Number string `json:"number"`
}

// TableName method for User struct
func (User) TableName() string { return "users" }

// TableName method for CreditCard struct
//func (CreditCard) TableName() string { return "creditcards" }
