package entity

import "time"

type Tickets struct {
	TicketID  int       `gorm:"primaryKey;autoIncrement" json:"ticketId"`
	EventID   int       `gorm:"not null" json:"eventId"`
	UserID    *int      `gorm:"null" json:"userId"`
	Status    string    `gorm:"type:enum('Available', 'Booked', 'Cancelled');not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

	Event Events `gorm:"foreignKey:EventID;references:EventID" json:"event"`
}
