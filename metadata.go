package main

import (
	"strings"
)

// Metadata contains meta data about node,
// extracted from name string.
type Metadata struct {
	Name      string
	Version   string
	OS        string
	GoVersion string
}

// UnknownMetadata represents metadata for failed to parse string.
var UnknownMetadata = &Metadata{"Unknown", "Unknown", "Unknown", "Unknown"}

// NewMetadata returns new Metadata object parsing string represtnation.
func NewMetadata(name string) *Metadata {
	fields := strings.Split(name, "/")
	if len(fields) != 4 {
		return UnknownMetadata
	}

	return &Metadata{
		Name:      fields[0],
		Version:   fields[1],
		OS:        fields[2],
		GoVersion: fields[3],
	}
}
