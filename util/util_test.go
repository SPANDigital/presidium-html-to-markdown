package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceTemplate(t *testing.T) {
	testCases := []struct {
		input    string
		args     []string
		expected string
	}{
		{
			input:    "{{% callout title=\"$1\" description=\"$2\" %}}",
			args:     []string{"Info", "Some text"},
			expected: "{{% callout title=\"Info\" description=\"Some text\" %}}",
		},
		{
			input:    "$$1",
			args:     []string{"test"},
			expected: "$test",
		},
		{
			input:    "$1$2$3",
			args:     []string{"1", "2"},
			expected: "12$3",
		},
	}

	for _, test := range testCases {
		actual := ReplaceTemplate(test.input, test.args)
		assert.Equal(t, test.expected, actual, test)
	}
}

func TestSlugify(t *testing.T) {
	testCases := map[string]string{
		"Hello World":         "hello-world",
		"How are you?":        "how-are-you",
		" multiple   space ":  "multiple-space",
		"$special#characters": "special-characters",
	}

	for test, expected := range testCases {
		actual := Slugify(test)
		assert.Equal(t, expected, actual, test)
	}
}

func TestIsURL(t *testing.T) {
	testCases := map[string]bool{
		"http://url":       true,
		"https://url":      true,
		"http://url/#test": true,
		"//url":            false,
		"file://url":       true,
		"/local/path":      false,
		"./local/path":     false,
		"local/path":       false,
		".":                false,
		"../":              false,
	}

	for test, expected := range testCases {
		actual := IsURL(test)
		assert.Equal(t, expected, actual, test)
	}
}

func TestIsExternalURL(t *testing.T) {
	baseDomain := "example.com"

	testCases := map[string]bool{
		"http://url":               true,
		"https://url":              true,
		"/url":                     false,
		"#test":                    false,
		"file://url":               true,
		"tel://url":                true,
		"mailto://hello@email.com": true,
		"/local/path":              false,
		"../local/path":            false,
		".":                        false,
	}

	for test, expected := range testCases {
		actual := IsExternalUrl(test, baseDomain)
		assert.Equal(t, expected, actual, test)
	}
}

func TestFileNameWithoutExt(t *testing.T) {
	testCases := map[string]string{
		"image.jpeg":     "image",
		"test-file.html": "test-file",
		"test.file.html": "test.file",
	}

	for test, expected := range testCases {
		actual := FileNameWithoutExt(test)
		assert.Equal(t, expected, actual, test)
	}
}
