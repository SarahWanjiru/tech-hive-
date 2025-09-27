package entity

import "time"

type User struct {
 	Id           uint      `gorm:"primaryKey;column:id;type:int;autoIncrement"`
 	Password     string    `gorm:"column:password;type:varchar(255);not null"`
 	Name         string    `gorm:"column:name;type:varchar(150);not null"`
 	Email        string    `gorm:"column:email;type:varchar(255);unique;not null"`
 	Role         string    `gorm:"column:role;type:varchar(50);default:customer;not null;check:role IN ('customer', 'admin')"`
 	IsActive     bool      `gorm:"column:is_active;type:boolean;default:true"`
 	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
 	Orders       []Order   `gorm:"ForeignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
 }

func (User) TableName() string {
	return "tb_user"
}
