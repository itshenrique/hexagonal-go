package domain

import "time"

type Session struct {
	ID           string
	Username     string
	LastModified *time.Time

	// Add other fields, when necessary
}
