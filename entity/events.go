package entity

import "time"

type Events struct {
	EventID    int       `gorm:"primaryKey;autoIncrement" json:"eventId"`
	EventName  string    `gorm:"type:varchar(255);not null;unique" json:"eventName"`
	Date       string    `json:"date"`
	ParsedDate time.Time `json:"-" gorm:"-"`
	Status     string    `gorm:"type:enum('Active', 'Ongoing', 'Completed');not null" json:"status"`
	Price      float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Capacity   int       `gorm:"not null;check:capacity >= 0" json:"capacity"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
