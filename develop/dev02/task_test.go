package main

import (
	"errors"
	"testing"
)

func TestUnpackStr(t *testing.T) {
	tests := []struct {
		testName string
		inputStr string
		required string
		err      error
	}{
		{testName: "1", inputStr: "a4bc2d5e", required: "aaaabccddddde", err: nil},
		{testName: "2", inputStr: "abcd", required: "abcd", err: nil},
		{testName: "3", inputStr: "45", required: "", err: InvalidInputError},
		{testName: "4", inputStr: "", required: "", err: nil},
	}
	for _, e := range tests {
		t.Run(e.testName, func(t *testing.T) {
			result, err := UnpackStr(e.inputStr)
			if result != e.required {
				t.Errorf("result: %v, required: %v", result, e.required)
			}
			if !errors.Is(err, e.err) {
				t.Errorf("result: %v, required: %v", err, e.err)
			}
		})
	}
}
