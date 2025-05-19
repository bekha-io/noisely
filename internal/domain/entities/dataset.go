package entities

import (
	"io"
	"net/http"
	"time"
)

// Dataset represents a dataset in the system.
// Contains metadata, parameters and data in JSON format,
// as well as references to the data schema being used.
type Dataset struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	SchemaID  string    `json:"schema_id" bson:"schema_id"`
	SchemaVer string    `json:"schema_version" bson:"schema_version"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	FileURL   string    `json:"file_url" bson:"file_url"`
}

func (d *Dataset) GetFile() (io.ReadCloser, error) {
	response, err := http.Get(d.FileURL)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
