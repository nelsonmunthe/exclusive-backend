package entity

import "time"

type Purchase struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Product_id  int       `json:"product_id"`
	Customer_id int       `json:"customer_id"`
	Discount    float64   `json:"discount"`
	Quantity    float64   `json:"quantity"`
	Total_price float64   `json:"total_price"`
	Shipping    float64   `json:"shipping"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
