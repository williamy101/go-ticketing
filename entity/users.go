package entity

import "time"

type Users struct {
	UserID    int       `gorm:"primaryKey;autoIncrement" json:"userId"`
	UserName  string    `gorm:"type:varchar(255);not null" json:"userName"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:enum('Admin','User');not null" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
