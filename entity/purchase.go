package entity

import "time"

type Purchase_line struct {
	ID                 int       `gorm:"primaryKey" json:"id"`
	Purchase_header_id int       `json:"purchase_header_id"`
	Product_id         int       `json:"product_id"`
	Discount           float64   `json:"discount"`
	Quantity           int       `json:"quantity"`
	Total_price        float64   `json:"total_price"`
	Shipping           float64   `json:"shipping"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
}

type Purchase_header struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Customer_id string    `json:"customer_id"`
	Total_price float64   `json:"total_price"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

type CreatePurchase struct {
	Purchase_line []Purchase_line `json:"purchase_line"`
	Total_price   int             `json:"total_price"`
}
