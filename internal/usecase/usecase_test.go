package usecase

import (
	"strings"
	"testing"
)

func Test_find_SubStr(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "first test",
			input:    "abcabcd",
			expected: "abcd",
		},
		{
			name:     "second test",
			input:    "aaaaaaaa",
			expected: "a",
		},
		{
			name:     "third test",
			input:    "qwertyq",
			expected: "qwerty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Find_SubStr(tc.input)
			if res != tc.expected {
				t.Errorf("Find_SubStr(%v) = %v, want %v", tc.input, res, tc.expected)
			}
		})
	}
}

func Test_Check_Email(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "first test",
			input:    "hi my name is blabla. and here is my Email:  blabla@gmail.com",
			expected: strings.Join([]string{"Email:  blabla@gmail.com"}, " "),
		},
		{
			name:     "second test",
			input:    "some text with Email: fake.email@gmail.com and some more text with Email:    another.fake.email@example.com   ",
			expected: strings.Join([]string{"Email: fake.email@gmail.com", "Email:    another.fake.email@example.com"}, " "),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Check_Email(tc.input)
			if strings.Join(res, " ") != tc.expected {
				t.Errorf("Check_Email(%v) = %v, want %v", tc.input, res, tc.expected)
			}
		})
	}

}

func Test_Find_self(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:  "first test",
			input: "int",
			expected: strings.Join([]string{
				"interface",
				}, ","),
			err: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := Find_Self(tc.input)
			if strings.Join(output, ",") != tc.expected || err != tc.err {
				t.Errorf("Find_Self(%v) = \n output: %v, err: %v,\n want %v", tc.input, output, err, tc.expected)

			}
		})
	}
}
