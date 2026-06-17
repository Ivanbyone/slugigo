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

/*
Helpers
*/

// isAllowed - helper function to check allowed ASCII symbols.
//
// Current allowed pattern: `[^a-zA-Z0-9\s\-_.]`
func isAllowed(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') ||
		char == '-' || char == '_' || char == '.'
}

// isSpace - helper function to check whitespaces, tabulations, etc.
func isSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

// isUppercase - helper function to check Uppercase ASCII symbols.
func isUppercase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

// trim - delete all around whitespaces
func trim(b []byte) []byte {
	return bytes.TrimSpace(b)
}

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

func (s Slugigo) SaveLeadingAndTrailingSeparator() Slugigo {
	s.flags |= flagSaveLeadingAndTrailingDash
	return s
}

// Separator
func (s Slugigo) Separator(sep string) Slugigo {
	s.separator = sep
	return s
}

// normalize
func (s Slugigo) normalize(buf []byte) []byte {
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

		if isSpace(char) {
			if !space {
				buf[w] = sep
				w++
				space = true
			}
			continue
		}

		if isAllowed(char) {

			if !hasNoLowercaseFlag && isUppercase(char) {
				char += 32
			}

			buf[w] = char
			w++
			space = false
		}
	}

	return buf[:w]
}

// preprocessing
func (s Slugigo) preprocessing(buffer []byte) []byte {
	buf := trim(buffer)
	return buf
}

// postprocessing
func (s Slugigo) postprocessing(buffer []byte) []byte {
	buf := s.removeLeadingAndTrailingSeparator(buffer)
	return buf
}

func (s Slugigo) removeLeadingAndTrailingSeparator(buf []byte) []byte {
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

// Build
func (s Slugigo) Build() string {
	buffer := s.preprocessing(s.text)
	buffer = s.normalize(buffer)
	buffer = s.postprocessing(buffer)
	return string(buffer)
}
