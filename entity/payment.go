package entity

import (
	"time"
)

type Payment struct {
 	Id            uint      `gorm:"primaryKey;column:id;type:int;autoIncrement"`
 	OrderId       uint      `gorm:"column:order_id;type:int;not null"`
 	Order         Order     `gorm:"ForeignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 	TransactionId string    `gorm:"column:transaction_id;type:varchar(255);unique"`
 	Status        string    `gorm:"column:status;type:varchar(50);default:pending;check:status IN ('pending', 'success', 'failed', 'cancelled')"`
 	PaidAt        time.Time `gorm:"column:paid_at;type:timestamp;default:CURRENT_TIMESTAMP"`
 }

func (Payment) TableName() string {
	return "tb_payment"
}