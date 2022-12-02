package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	expected := 1
	actual := 1

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
