package main

import (
	"math/rand"
	"testing"
	"time"
)

func makeTest(maxNum int) (slice []int, missNum int) {
	rand.Seed(time.Now().Unix())
	missNum = rand.Intn(maxNum)

	slice = make([]int, maxNum+1, maxNum+1)
	for i := 0; i <= maxNum; i++ {
		slice[i] = i
	}

	slice = append(slice[:missNum], slice[missNum+1:]...)
	return
}

func Benchmark_missingNumber(b *testing.B) {
	slice, _ := makeTest(100000000)
	b.Run("make() and two loops", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			missingNumber(slice)
		}
	})

	b.Run("Trapezoid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Trapezoid(slice)
		}
	})

	b.Run("Complement", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Complement(slice)
		}
	})

	b.Run("Cathy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Cathy(slice)
		}
	})

	b.Run("Lisa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Lisa(slice)
		}
	})

}

func Test_missingNumber(t *testing.T) {
	slice, missNum := makeTest(4)

	assertEqual := func(name string, got, ans int) {
		t.Helper()
		if got != ans {
			t.Errorf("[%v] got %v, expect %v",
				name,
				got,
				ans,
			)
		}
	}

	t.Run("make() and two loops", func(t *testing.T) {
		got := missingNumber(slice)
		assertEqual("Naive", got, missNum)
	})

	t.Run("Trapezoid", func(t *testing.T) {
		got := Trapezoid(slice)
		assertEqual("Trapezoid", got, missNum)
	})

	t.Run("Complement", func(t *testing.T) {
		got := Complement(slice)
		assertEqual("Complement", got, missNum)
	})
}
