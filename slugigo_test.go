/*
Copyright (c) 2026 Ivan Boyko.

This source code is licensed under the MIT license found in the
LICENSE file in the root directory of this source tree.
*/

package slugigo

import "testing"

func TestIsAllowed(t *testing.T) {
	cases := []struct {
		provided byte
		expected bool
	}{
		{'a', true},
		{'o', true},
		{'z', true},
		{'0', true},
		{'5', true},
		{'9', true},
		{'.', true},
		{'-', true},
		{'_', true},
		{'A', true},
		{'Z', true},
		{'?', false},
		{'!', false},
		{',', false},
	}

	for _, c := range cases {
		actual := isAllowed(c.provided)
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %t, actual: %v", c.expected, actual)
		}
	}
}

func TestIsSpace(t *testing.T) {
	cases := []struct {
		provided byte
		expected bool
	}{
		{' ', true},
		{'\n', true},
		{'\t', true},
		{'\r', true},
		{'a', false},
		{'.', false},
		{'_', false},
	}

	for _, c := range cases {
		actual := isSpace(c.provided)
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %t, actual: %v", c.expected, actual)
		}
	}
}

func TestIsUppercase(t *testing.T) {
	cases := []struct {
		provided byte
		expected bool
	}{
		{'A', true},
		{'O', true},
		{'Z', true},
		{'a', false},
		{'z', false},
		{'-', false},
		{'.', false},
		{'_', false},
	}

	for _, c := range cases {
		actual := isUppercase(c.provided)
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %t, actual: %v", c.expected, actual)
		}
	}
}

func TestDefaultSlugCreation(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"", ""},
		{" hello", "hello"},
		{"hello slug.   ", "hello-slug."},
		{"   Hello, Slugigo!      ", "hello-slugigo"},
		{"this is a text for slug", "this-is-a-text-for-slug"},
		{"this-is a text._For slug!,", "this-is-a-text._for-slug"},
		{".  this-is a text._For slug!,   ?", ".-this-is-a-text._for-slug"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}

func TestCustomSeparator(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"Hello, Slug!", "hello+slug"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			Separator("+").
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}

func TestNoLowercase(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"Hello, Slug!", "Hello-Slug"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			NoLowercase().
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}

func TestMaxLength(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"Hello, Slug!", "hello-sl"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			MaxLength(8).
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}

func TestRemoveLeadingAndTrailingDashes(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"---Hello Slug---", "---hello-slug---"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			SaveLeadingAndTrailingDash().
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}
