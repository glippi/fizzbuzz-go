package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("returns 1", func(t *testing.T) {
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
	t.Run("returns 'FizzBuzz' when argument is divisible by 3 and 5", func(t *testing.T) {
		got := FizzBuzz(15)
		want := "FizzBuzz"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
	t.Run("returns 2", func(t *testing.T) {
		got := FizzBuzz(2)
		want := "2"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
