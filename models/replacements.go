package models

import "time"

// Replacement represents an element in the
// replacement step of the protocol analyse pipeline.
type Replacement struct {
	ID            string    `gorm:"primary_key"`
	SearchString  string    `gorm:"not null"`
	ReplaceString string    `gorm:"not null"`
	Created       time.Time `gorm:"not null"`
}
