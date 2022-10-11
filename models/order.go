package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null;type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
