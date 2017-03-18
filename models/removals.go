package models

import "time"

// Removal represents an element in the
// removal step of the protocol analyse pipeline.
type Removal struct {
	ID       string    `gorm:"primary_key"`
	StartTag string    `gorm:"not null"`
	EndTag   string    `gorm:"not null"`
	Created  time.Time `gorm:"not null"`
}
