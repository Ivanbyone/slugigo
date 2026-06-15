package slugigo

import "testing"

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
