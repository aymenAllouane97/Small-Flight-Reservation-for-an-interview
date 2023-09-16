package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string   `gorm:"not null"`
	Email    string   `gorm:"not null;uniqueIndex;matches=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
	Password string   `gorm:"not null"`
	Flights  []Flight `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (c *Company) ValidatePasswordLength() bool {
	return len(c.Password) >= 8
}
