package fizzBuzz

import (
	"strconv"
	"strings"
)

func FizzBuzzIt(numbers []int) string {
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
