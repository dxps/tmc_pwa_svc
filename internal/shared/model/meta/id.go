package meta

import (
	"log/slog"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Id is a unique identifier.
type Id string

func NewId() Id {
	id, err := gonanoid.New(10)
	if err != nil {
		slog.Error("Failed to generate id.", "error", err)
	}
	return Id(id)
}

func (id Id) String() string {
	return string(id)
}
