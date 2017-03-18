package models

import "time"

// Removal represents an element in the removal
// part of the protocol analysis pipeline.
type Removal struct {
	ID       string    `gorm:"primary_key"`
	Created  time.Time `gorm:"not null"`
	StartTag string    `gorm:"not null"`
	EndTag   string    `gorm:"not null"`
}
