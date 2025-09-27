package entity

import (
	"time"
	"github.com/google/uuid"
)

type Product struct {
 	Id          uint          `gorm:"primaryKey;column:id;type:int;autoIncrement"`
 	ProductId   uuid.UUID     `gorm:"column:product_id;type:varchar(36);unique;not null"`
 	Name        string        `gorm:"index;column:name;type:varchar(100);not null"`
 	Description string        `gorm:"column:description;type:text"`
 	Price       float64       `gorm:"column:price;type:decimal(10,2);not null"`
 	Stock       int32         `gorm:"column:stock;type:int;default:0;not null"`
 	ImageUrl    string        `gorm:"column:image_url;type:text"`
 	CreatedAt   time.Time     `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
 	OrderItems  []OrderItem   `gorm:"ForeignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 }

func (Product) TableName() string {
	return "tb_product"
}
