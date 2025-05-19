package entities

import (
	"context"
	"io"
)

// Exporter is an interface for exporting datasets.
type Exporter interface {
	// Export exports the dataset in the format defined by the exporter.
	// Returns an io.ReadCloser for reading the exported data.
	Export(ctx context.Context, dataset Dataset) (io.ReadCloser, error)
}
