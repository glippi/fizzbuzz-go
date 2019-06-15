package main

import "testing"

func TestDivisibleBy(t *testing.T) {
	t.Run("returns true if the modulo between a and b is not 0", func(t *testing.T) {
		got := isDivisibleByWithNorest(10, 3)
		want := false

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
	t.Run("returns true if the modulo between a and b is 0", func(t *testing.T) {
		got := isDivisibleByWithNorest(15, 3)
		want := true

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}

func TestFizzBuzz(t *testing.T) {
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
	t.Run("returns the passed argument if is not divisible by 3 or 5", func(t *testing.T) {
		tt := []struct{
			input int
			output string
		}{
			{
				input: 2,
				output: "2",
			},
			{
				input: 4,
				output: "4",
			},
		} 


		for _, tc  := range tt {
			got := FizzBuzz(tc.input)
			want := tc.output

			if got != want {
				t.Errorf("got '%s' want '%s'", got, want)
			}

		}


	})
}
