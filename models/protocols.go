package models

import "time"

// Category denotes the type of a protocol
// item. For example, being from one of the
// Freitagsrunde's internal meetings.
const CategoryFreitagssitzung = iota

// Status represents the phase an uploaded
// protocol is currently in. From 'in review'
// to 'published'.
const (
	StatusInReview = iota
	StatusPublished
)

// Protocol represents an internal meeting protocol
// to be reviewed and subsequently published by use
// of protokollamt.
type Protocol struct {
	ID              string    `gorm:"primary_key"`
	UploadDate      time.Time `gorm:"not null"`
	MeetingDate     time.Time `gorm:"not null"`
	Category        int       `gorm:"index;not null"`
	InternalVersion string    `gorm:"not null"`
	PublicVersion   string    `gorm:"not null"`
	Status          int       `gorm:"not null"`
}
