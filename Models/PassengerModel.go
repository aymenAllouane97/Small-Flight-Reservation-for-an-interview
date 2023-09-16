package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	Name         string        `gorm:"not null"`
	Reservations []Reservation `gorm:"foreignKey:PassengerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
