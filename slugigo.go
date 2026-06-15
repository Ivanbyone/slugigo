/*
Copyright (c) 2026 Ivan Boyko.

This source code is licensed under the MIT license found in the
LICENSE file in the root directory of this source tree.
*/

package slugigo

import (
	"bytes"
)

const (
	flagNoLowercase uint8 = 1 << iota
	flagMaxLength
	flagSaveLeadingAndTrailingDash
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

// NoLowercase
func (s Slugigo) NoLowercase() Slugigo {
	s.flags |= flagNoLowercase
	return s
}

// MaxLength
func (s Slugigo) MaxLength(length int) Slugigo {
	s.max = length
	s.flags |= flagMaxLength
	return s
}

func (s Slugigo) SaveLeadingAndTrailingDash() Slugigo {
	s.flags |= flagSaveLeadingAndTrailingDash
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
	hasNoLowercaseFlag := s.flags&flagNoLowercase != 0
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

			if !hasNoLowercaseFlag && char >= 'A' && char <= 'Z' {
				char += 32
			}

			buf[w] = char
			w++
			space = false
		}
	}

	return buf[:w]
}

func (s Slugigo) removeLeadingAndTrailingDash(buf []byte) []byte {
	hasSaveLeadingAndTrailingDashFlag := s.flags&flagSaveLeadingAndTrailingDash != 0
	sep := s.separator[0]

	for !hasSaveLeadingAndTrailingDashFlag && len(buf) > 0 && buf[0] == sep {
		buf = buf[1:]
	}
	for !hasSaveLeadingAndTrailingDashFlag && len(buf) > 0 && buf[len(buf)-1] == sep {
		buf = buf[:len(buf)-1]
	}

	return buf
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
	// 3. Process leading and trailing dashes
	buffer = s.removeLeadingAndTrailingDash(buffer)
	return string(buffer)
}
