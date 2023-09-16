package models

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	gorm.Model
	CompanyID          uint          `gorm:"not null" json:"company_id"`
	DepartureAirportID uint          `gorm:"not null" json:"departure_airport_id"`
	ArrivalAirportID   uint          `gorm:"not null" json:"arrival_airport_id"`
	DepartureDate      time.Time     `gorm:"not null" json:"departure_date"`
	ArrivalDate        time.Time     `gorm:"not null" json:"arrival_date"`
	DepartureHour      time.Time     `gorm:"not null" json:"departure_hour"`
	ArrivalHour        time.Time     `gorm:"not null" json:"arrival_hour"`
	Reservations       []Reservation `gorm:"foreignKey:FlightID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Airports           []Airport     `gorm:"many2many:Stopover;"`
}

func (f *Flight) TableName() string {
	return "flights"
}
