Algoritm:

Step 1: Define a function that accepts a numbers(num) type is int, till then need to print the series.
Step 2: Take two initial numbers for the series, i.e., 0 and 1.
Step 3: Start a true for loop and declare a third variable to store previous two values.
Step 4: Print the summation of two numbers until the summation is lesser than the given num.






Code:

package main

import "fmt"

func printFibonacciSeries(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Series is: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

func main() {
	printFibonacciSeries(10)

}
