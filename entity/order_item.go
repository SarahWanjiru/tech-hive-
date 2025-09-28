package entity

import (
	"time"
	"github.com/google/uuid"
)

type OrderItem struct {
  	Id        uint      `gorm:"primaryKey;column:id;type:int;autoIncrement"`
  	OrderId   uint      `gorm:"column:order_id;type:int;not null"`
  	Order     Order     `gorm:"ForeignKey:OrderId;References:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
  	ProductId uuid.UUID `gorm:"column:product_id;type:varchar(36);not null"`
  	Product   Product   `gorm:"ForeignKey:ProductId;References:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
  	Quantity  int32     `gorm:"column:quantity;type:int;not null;check:quantity > 0"`
  	Price     float64   `gorm:"column:price;type:decimal(10,2);not null;check:price >= 0"`
  	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
   }

func (OrderItem) TableName() string {
	return "tb_order_item"
}