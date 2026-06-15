package slugigo

import "testing"

func TestDefaultSlugCreation(t *testing.T) {
	cases := []struct {
		provided string
		expected string
	}{
		{"", ""},
		{" helloSlug", "helloSlug"},
		{"hello slug.   ", "hello slug."},
		{"   Hello, Slug!         ", "Hello, Slug!"},
		// {"   Hello, Slugigo!      ", "Hello-Slugigo"},
		// {"this is a text for slug", "this-is-a-text-for-slug"},
		// {"this-is a text._For slug!,", "this-is-a-text._For-slug"},
		// {".  this-is a text._For slug!,   ?", ".-this-is-a-text._For-slug-"},
	}

	for _, c := range cases {
		actual := Slug(c.provided).Build()
		if actual != c.expected {
			t.Errorf("Test failed! Expected: %s, actual: %s", c.expected, actual)
		}
	}
}
