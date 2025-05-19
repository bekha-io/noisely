package entities

import "time"

// Schema represents a data schema in the system.
// Contains metadata and schema specification in JSON format.
type Schema struct {
	ID        string    `json:"id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Spec      []byte    `json:"spec"`
}
