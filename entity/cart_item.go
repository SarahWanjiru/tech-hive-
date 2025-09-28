package entity

import "time"

type CartItem struct {
 	Id        uint      `gorm:"primaryKey;column:id;type:int;autoIncrement"`
 	CartId    uint      `gorm:"column:cart_id;type:int;not null"`
 	Cart      Cart      `gorm:"ForeignKey:CartId;References:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 	ProductId string    `gorm:"column:product_id;type:varchar(36);not null"`
 	Product   Product   `gorm:"ForeignKey:ProductId;References:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 	Quantity  int32     `gorm:"column:quantity;type:int;not null;check:quantity > 0"`
 	Price     float64   `gorm:"column:price;type:decimal(10,2);not null;check:price >= 0"`
 	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (CartItem) TableName() string {
	return "tb_cart_item"
}