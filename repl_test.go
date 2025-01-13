package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hell0 World ",
			expected: []string{"hell0", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
		}

		for i := range actual {
			word := actual[i]

			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
