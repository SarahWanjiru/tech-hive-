package entity

import "time"

type Cart struct {
	Id        uint      `gorm:"primaryKey;column:id;type:int;autoIncrement"`
	UserId    uint      `gorm:"column:user_id;type:int;not null"`
	User      User      `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	CartItems []CartItem `gorm:"ForeignKey:CartId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Cart) TableName() string {
	return "tb_cart"
}