package leapYear

import "testing"

func TestLeapYear(t *testing.T) {
	t.Run("year 1997 is not leap as not divisible by 4", func(t *testing.T) {
		want := false
		got := isLeap(1997)
		if got != want {
			t.Errorf("Year 1997 is not a leap year")
		}
	})
	t.Run("year 1996 is a leap year as divisible by 4", func(t *testing.T) {
		want := true
		got := isLeap(1996)
		if got != want {
			t.Errorf("Year 1996 is a leap year")
		}
	})
	t.Run("year 1600 is a leap year as divisible by 400", func(t *testing.T) {
		want := true
		got := isLeap(1600)
		if got != want {
			t.Errorf("Year 1600 is a leap year")
		}
	})
	t.Run("year 1800 is not a leap year as is divisible by 100 but not 400", func(t *testing.T) {
		want := false
		got := isLeap(1800)
		if got != want {
			t.Errorf("Year 1800 is a not leap year")
		}
	})
}

func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
