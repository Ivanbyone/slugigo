package slugigo

import (
	"bytes"
)

const (
	flagLowercase uint8 = 1 << iota
	flagMaxLength
)

// Slugigo
type Slugigo struct {
	flags     uint8
	max       int
	separator string
	text      []byte
}

// Slug returns a new Slugigo instance initialized with the provided string.
//
// It provides a Fluent API for configuring the slug generator through method chaining.
// Use Build() to generate the final slug after setting desired options.
func Slug(text string) Slugigo {
	b := []byte(text)
	return Slugigo{text: b, separator: "-"}
}

// Lowercase
func (s Slugigo) Lowercase() Slugigo {
	s.flags |= flagLowercase
	return s
}

// MaxLength
func (s Slugigo) MaxLength(length int) Slugigo {
	s.max = length
	s.flags |= flagMaxLength
	return s
}

// Separator
func (s Slugigo) Separator(sep string) Slugigo {
	s.separator = sep
	return s
}

// process
//
// Allowed pattern: `[^a-zA-Z0-9\s\-_.]`
func (s Slugigo) process(buf []byte) []byte {
	w := 0
	sep := s.separator[0]
	space := false

	// Check flags
	hasLowercaseFlag := s.flags&flagLowercase != 0
	hasMaxLengthFlag := s.flags&flagMaxLength != 0

	for i := range buf {
		char := buf[i]

		if hasMaxLengthFlag && w >= s.max {
			break
		}

		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			if !space {
				buf[w] = sep
				w++
				space = true
			}
			continue
		}

		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '-' || char == '_' || char == '.' {

			if hasLowercaseFlag && char >= 'A' && char <= 'Z' {
				char += 32
			}

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
	buffer = s.process(buffer)
	return string(buffer)
}
