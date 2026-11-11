package main

import (
	"testing"
)

func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Pikachu CHarizod Balbazor",
			expected: []string{"pikachu","charizod","balbazor"},
		},
		{
			input: "  leading white space",
			expected: []string{"leading", "white", "space"},
		},
		{
			input: "Trailing white space  ",
			expected: []string{"trailing", "white", "space"},
		},
	}
		
	//now let's check the test cases
	
	for _, c := range cases{
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("expected %d got %d words", len(c.expected), len(actual))
			continue
		}

		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord{
				t.Errorf("expected %s got %s", expectedWord, word)
				continue
			}
		}
	}
}
