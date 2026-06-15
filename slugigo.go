package slugigo

import (
	"bytes"
)

type Slugigo struct {
	text      []byte
	separator string
}

// Slug returns a new Slugigo instance initialized with the provided string.
//
// It provides a Fluent API for configuring the slug generator through method chaining.
// Use Build() to generate the final slug after setting desired options.
func Slug(text string) Slugigo {
	b := []byte(text)
	return Slugigo{text: b, separator: "-"}
}

// Separator
func (s Slugigo) Separator(sep string) Slugigo {
	s.separator = sep
	return s
}

// clean
//
// Allowed pattern: `[^a-zA-Z0-9\s\-_.]`
func (s Slugigo) clean(buf []byte) []byte {
	w := 0
	sep := s.separator[0]
	space := false

	for i := range buf {
		char := buf[i]

		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			if !space {
				buf[w] = sep
				w++
				space = true
			}
		}

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') ||
			char == '-' || char == '_' || char == '.' {
			buf[w] = char
			w++
			space = false
		}
	}

	return buf[:w]
}

// trim
func (s Slugigo) trim(b []byte) []byte {
	return bytes.TrimSpace(b)
}

// Build
func (s Slugigo) Build() string {
	// 1. Trim (mandatory operation)
	buffer := s.trim(s.text)

	// 2. Remove Special Symbols
	buffer = s.clean(buffer)

	return string(buffer)
}
