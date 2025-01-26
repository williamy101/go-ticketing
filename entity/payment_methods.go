package entity

import "time"

type PaymentMethods struct {
	MethodID   int       `gorm:"primaryKey;autoIncrement" json:"methodId"`
	MethodName string    `gorm:"type:varchar(100);not null;unique" json:"methodName"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
