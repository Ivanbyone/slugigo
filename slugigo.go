package slugigo

import (
	"bytes"
)

type Slugigo struct {
	text []byte
}

// Slug returns a new Slugigo instance initialized with the provided string.
// 
// It provides a Fluent API for configuring the slug generator through method chaining.
// Use Build() to generate the final slug after setting desired options.
func Slug(text string) Slugigo {
	b := []byte(text)
	return Slugigo{text: b}
}

// trim 
func (s Slugigo) trim(b []byte) []byte {
	return bytes.TrimSpace(b)
}

// Build
func (s Slugigo) Build() string {
	// 1. Trim (mandatory operation)
	b := s.trim(s.text)

	return string(b)
}
