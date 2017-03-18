package models

import "time"

// Replacement represents an element in the replacement
// part of the protocol analysis pipeline.
type Replacement struct {
	ID            string    `gorm:"primary_key"`
	Created       time.Time `gorm:"not null"`
	SearchString  string    `gorm:"not null"`
	ReplaceString string    `gorm:"not null"`
}
