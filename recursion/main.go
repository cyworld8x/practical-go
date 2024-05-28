package recursion

import (
	"fmt"
)

// Go supports recursive functions
// A recursive function is a function that calls itself
// Recursion is a common mathematical and programming concept. It means that a function calls itself. This has the benefit of meaning that you can loop through data to reach a result.
// The developer should be very careful with recursion as it can be quite easy to slip into writing a function which never terminates, or one that uses excess amounts of memory or processor power. However, when written correctly, recursion can be a very efficient and mathematically-elegant approach to programming.
// In this example, we will learn how to write a recursive function to calculate the factorial of a number.
// The factorial of a positive integer n is the product of all positive integers less than or equal to n.
// The factorial of 0 is 1.
// The factorial of 1 is 1.
// The factorial of 2 is 2.
// The factorial of 3 is 6.
// The factorial of 4 is 24.

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Run() {
	fmt.Println("[Recursion] ----------> Start Run Fibonacci(10) <----------")
	fib := Fibonacci(10)
	fmt.Println(fib)
}
