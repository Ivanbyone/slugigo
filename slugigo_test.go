package slugigo

import "testing"

func TestDefaultSlugCreation(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"", ""},
		{" helloSlug", "helloSlug"},
		{"hello slug.   ", "hello-slug."},
		{"   Hello, Slugigo!      ", "Hello-Slugigo"},
		{"this is a text for slug", "this-is-a-text-for-slug"},
		{"this-is a text._For slug!,", "this-is-a-text._For-slug"},
		{".  this-is a text._For slug!,   ?", ".-this-is-a-text._For-slug-"},
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
		{"Hello, Slug!", "Hello+Slug"},
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

func TestLowercase(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"Hello, Slug!", "hello-slug"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).
			Lowercase().
			Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}
