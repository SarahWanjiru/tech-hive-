package entity

import (
	"time"
)

type Order struct {
 	Id        uint        `gorm:"primaryKey;column:id;type:int;autoIncrement"`
 	UserId    uint        `gorm:"column:user_id;type:int;not null"`
 	User      User        `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 	Total     float64     `gorm:"column:total;type:decimal(10,2);not null;check:total >= 0"`
 	Status    string      `gorm:"column:status;type:varchar(50);default:pending;check:status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled')"`
 	CreatedAt time.Time   `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
 	OrderItems []OrderItem `gorm:"ForeignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 	Payments   []Payment   `gorm:"ForeignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 }

func (Order) TableName() string {
	return "tb_order"
}