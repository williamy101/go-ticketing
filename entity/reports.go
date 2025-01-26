package entity

type ReportSummary struct {
	TotalTicketsSold int     `json:"totalTicketsSold" gorm:"column:totalTicketsSold"`
	TotalRevenue     float64 `json:"totalRevenue" gorm:"column:totalRevenue"`
}

type EventReport struct {
	EventName   string  `json:"eventName" gorm:"column:eventName"`
	TicketsSold int     `json:"ticketsSold" gorm:"column:ticketsSold"`
	Revenue     float64 `json:"revenue" gorm:"column:revenue"`
}
