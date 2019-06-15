package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("returns argument if not divisible by 3 and/or 5", func(t *testing.T) {
		got := FizzBuzz(1)
		want := "1"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("returns 'Fizz' when argument is divisible by 3", func(t *testing.T) {
		got := FizzBuzz(3)
		want := "Fizz"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("returns 'Buzz' when argument is divisible by 5", func(t *testing.T) {
		got := FizzBuzz(5)
		want := "Buzz"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
