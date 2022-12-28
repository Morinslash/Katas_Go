package fizzBuzz

import (
	"strconv"
	"strings"
	"testing"
)

func TestFizzBuzzer(t *testing.T) {
	t.Run("given number 1 returns 1 as string", func(t *testing.T) {
		want := "1"
		got := fizzBuzzIt([]int{1})
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
	})
	t.Run("given number from 1 to 2 returns string 1\n2\n", func(t *testing.T) {
		want := "1\n2"
		got := fizzBuzzIt([]int{1, 2})
		if got != want {
			t.Errorf("expected %q got %q", want, got)
		}
	})
}

func fizzBuzzIt(numbers []int) string {
	result := ""
	for _, number := range numbers {
		result += (strconv.Itoa(number) + "\n")
	}
	return strings.Trim(result, "\n")
}
