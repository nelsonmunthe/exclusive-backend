package entity

import "time"

type Product struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category_id int     `json:"category_id"`
	Image_Url   string  `json:"image_Url"`
	Rating      float64 `json:"rating"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Total       int     `json:"total"`
	Flash_Sell  bool    `json:"flash_sell"`
	Images      string  `json:"images"`
	Created_at  time.Time
	Updated_at  time.Time
}
