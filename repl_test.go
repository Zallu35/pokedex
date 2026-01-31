package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Bulbasaur Ivysaur Venasaur ",
			expected: []string{"bulbasaur", "ivysaur", "venasaur"},
		},
		{
			input:    "thisonlylooks,for whitespace:soweird punctuation_won'taffect it",
			expected: []string{"thisonlylooks,for", "whitespace:soweird", "punctuation_won'taffect", "it"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual differs from length of expected, a: %v - e: %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %d does not match! Expected: %v, Actual: %v", i+1, expectedWord, word)
			}
		}
	}
}
