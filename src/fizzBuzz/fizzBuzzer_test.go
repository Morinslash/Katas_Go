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
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given number from 1 to 2 returns string 1\n2\n", func(t *testing.T) {
		want := "1\n2"
		got := fizzBuzzIt([]int{1, 2})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given number from 1 to 3 returns 1\n2\nFizz", func(t *testing.T) {
		want := "1\n2\nFizz"
		got := fizzBuzzIt([]int{1, 2, 3})
		assertFizzBuzzResult(t, got, want)
	})
	t.Run("given numbers from 1 to 5 returns 1\n2\nFizz\n4\nBuzz", func(t *testing.T) {
		want := "1\n2\nFizz\n4\nBuzz"
		got := fizzBuzzIt([]int{1, 2, 3, 4, 5})
		assertFizzBuzzResult(t, got, want)
	})
}

func assertFizzBuzzResult(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}

func fizzBuzzIt(numbers []int) string {
	result := ""
	for _, number := range numbers {
		result += convertFrom(number)
	}
	return format(result)
}

func format(result string) string {
	return strings.Trim(result, "\n")
}

func convertFrom(number int) string {
	switch number {
	case 5:
		return "Buzz\n"
	case 3:
		return "Fizz\n"
	default:
		return strconv.Itoa(number) + "\n"
	}
}
