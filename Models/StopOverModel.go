package models

import (
	"time"
)

type Stopover struct {
	FlightID  uint `gorm:"primaryKey"`
	AirportID uint `gorm:"primaryKey"`
	Arrival   time.Time
	Departure time.Time
}
