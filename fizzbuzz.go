package main

import "strconv"

func isDivisibleByWithNorest(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false

}

func FizzBuzz(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "FizzBuzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else if number%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}
