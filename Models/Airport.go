package models

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name             string `gorm:"not null"`
	CityID           uint
	DepartureFlights []Flight `gorm:"foreignKey:DepartureAirportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ArrivalFlights   []Flight `gorm:"foreignKey:ArrivalAirportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
