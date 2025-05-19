package entities

import "time"

// Dataset represents a dataset in the system.
// Contains metadata, parameters and data in JSON format,
// as well as references to the data schema being used.
type Dataset struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	SchemaID  string    `json:"schema_id" bson:"schema_id"`
	SchemaVer string    `json:"schema_version" bson:"schema_version"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Params    []byte    `json:"params" bson:"params"`
	Meta      []byte    `json:"meta" bson:"meta"`
	Data      []byte    `json:"data" bson:"data"`
}
