package operator // same package

import (
	"testing"
)

func Test_Add(t *testing.T) {
	if got := Add(1, 1); got != 2 {
		t.Errorf("Add(%v, %v) = %v, want %v", 1, 2, got, 2)
	}
}

func Test_subtract(t *testing.T) {
	if got := subtract(1, 1); got != 0 {
		t.Errorf("subtract(%v, %v) = %v, want %v", 1, 2, got, 2)
	}
}
