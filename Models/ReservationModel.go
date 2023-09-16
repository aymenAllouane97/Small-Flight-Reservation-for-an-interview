package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ClientID    uint `gorm:"not null" json:"client_id"`
	PassengerID uint `gorm:"not null" json:"passenger_id"`
	FlightID    uint `gorm:"not null" json:"flight_id"`
}

func (r *Reservation) TableName() string {
	return "reservations"
}
