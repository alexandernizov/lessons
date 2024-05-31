package packer_test

import (
	"errors"
	"testing"

	"github.com/alexandernizov/lessons/2/packer"
)

func TestStringUnpack(t *testing.T) {
	testTable := []struct {
		input    string
		expected string
		err      error
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{input: "abcd", expected: "abcd", err: nil},
		{input: "3abc", expected: "", err: packer.ErrIncorrectInput},
		{input: "45", expected: "", err: packer.ErrIncorrectInput},
		{input: "aaa10b", expected: "", err: packer.ErrIncorrectInput},
		{input: "aaa0b", expected: "aab", err: nil},
		{input: "", expected: "", err: nil},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc", err: nil},
	}

	for _, testCase := range testTable {
		res, err := packer.UnpackString(testCase.input)

		if testCase.err != nil && err == nil {
			t.Errorf("expected error: %s, got: nil", testCase.err.Error())
		}
		if testCase.err != nil && err != nil {
			if errors.Is(err, testCase.err) {
				continue
			}
		}
		if err != nil {
			t.Errorf("unhandled error: %s", err.Error())
		}

		if res != testCase.expected {
			t.Errorf("incorrect result. Expected: %q, got %q", testCase.expected, res)
		}
	}
}
