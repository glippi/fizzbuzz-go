package main

import "strconv"

func isDivisibleByWithNorest(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false

}

func FizzBuzz(number int) string {
	if isDivisibleByWithNorest(number, 3) && isDivisibleByWithNorest(number, 5) {
		return "FizzBuzz"
	} else if isDivisibleByWithNorest(number, 3) {
		return "Fizz"
	} else if isDivisibleByWithNorest(number, 5) {
		return "Buzz"
	} else {
		return strconv.Itoa(number)
	}
}
