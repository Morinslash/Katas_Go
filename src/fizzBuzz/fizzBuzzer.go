package fizzBuzz

import (
	"strconv"
	"strings"
)

func BuzzIt(numbers []int) string {
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
	if number%15 == 0 {
		return "FizzBuzz\n"
	}
	if number%5 == 0 {
		return "Buzz\n"
	}
	if number%3 == 0 {
		return "Fizz\n"
	}
	return strconv.Itoa(number) + "\n"

}
