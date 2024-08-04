package models

type Products struct {
	ID          uint `gorm:"primary_key" json:"id" query:"id" params:"id"`
	Name        string
	Description string
	Price       float64
}
