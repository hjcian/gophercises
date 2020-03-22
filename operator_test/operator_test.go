package operator_test //different package

import (
	"testing"

	"github.com/hjcian/gophercises/operator"
)

func Test_Add(t *testing.T) {
	if got := operator.Add(1, 1); got != 2 {
		t.Errorf("Add(%v, %v) = %v, want %v", 1, 2, got, 2)
	}
}

func Test_subtract(t *testing.T) {
	// get error here, because operator.subtract can't be exported
	if got := operator.subtract(1, 1); got != 2 {
		t.Errorf("Add(%v, %v) = %v, want %v", 1, 2, got, 2)
	}
}
