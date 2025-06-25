package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world  my  love  ",
			expected: []string{"hello", "world", "my", "love"},
		},
		{
			input:    "lovE",
			expected: []string{"love"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			excpectedWord := c.expected[i]
			if word != excpectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
