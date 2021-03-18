package lib

import "testing"

func AssertCorrect(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Fatalf("incorrect: expected %v but got %v", expected, actual)
	}
}
