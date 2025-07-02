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
			input:    "First TeSt",
			expected: []string{"first", "test"},
		},
		{
			input:    "Second   teSt",
			expected: []string{"second", "test"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match \n cleanInput(%q) == %q, want %q", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words don't match \n cleanInput(%q) == %q, want %q", c.input, word, expectedWord)
			}
		}
	}
}
