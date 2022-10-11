package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"not null;unique;type:varchar(50)" json:"itemCode"`
	description string `gorm:"not null" json:"descriprion"`
	quantity    uint   `gorm:"not null" json:"quantity"`
	OrderID     uint
}
