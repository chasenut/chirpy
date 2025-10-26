package main

import "testing"

func TestReplaceAllWordsCaseInsensitive(t *testing.T) {
	cases := []struct{
		input 		string
		filters 	[]string
		replacement	string
		expected 	string
	} {
		{
			input: "hello, world",
			filters: []string{},
			replacement: "****",
			expected: "hello, world",
		},
		{
			input: "hello, world",
			filters: []string{
				"world",
			},
			replacement: "****",
			expected: "hello, ****",
		},
		{
			input: "Hi there, how are you doing tHere fella?",
			filters: []string{
				"there",
			},
			replacement: "****",
			expected: "Hi there, how are you doing **** fella?",
		},
		{
			input: "mAYo KetCHup mUSTarD",
			filters: []string{
				"mayo", 
				"musta",
			},
			replacement: "****",
			expected: "**** KetCHup mUSTarD",
		},
	}

	for _, c := range cases {
		actual := replaceAllWordsCaseInsensitive(c.input, c.filters, c.replacement)
		if c.expected != actual {
			t.Errorf("ReplaceAllWordsCaseInsensitive generated: %s\n\nWant: %s", actual, c.expected)
			t.Fail()
		}
	}
}
