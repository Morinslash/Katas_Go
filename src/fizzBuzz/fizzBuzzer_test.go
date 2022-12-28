package fizzBuzz

import (
	"testing"
)

func TestFizzBuzzer(t *testing.T) {
	t.Run("given number 1 returns 1 as string", func(t *testing.T) {
		want := "1"
		got := BuzzIt([]int{1})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given number from 1 to 2 returns string 1\n2\n", func(t *testing.T) {
		want := "1\n2"
		got := BuzzIt([]int{1, 2})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given number from 1 to 3 returns 1\n2\nFizz", func(t *testing.T) {
		want := "1\n2\nFizz"
		got := BuzzIt([]int{1, 2, 3})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given numbers from 1 to 5 returns 1\n2\nFizz\n4\nBuzz", func(t *testing.T) {
		want := "1\n2\nFizz\n4\nBuzz"
		got := BuzzIt([]int{1, 2, 3, 4, 5})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given numbers from 11 to 15 returns 11\nFizz\n13\n14\nFizzBuzz", func(t *testing.T) {
		want := "11\nFizz\n13\n14\nFizzBuzz"
		got := BuzzIt([]int{11, 12, 13, 14, 15})
		assertFizzBuzzResult(t, got, want)
	})
}

func assertFizzBuzzResult(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
