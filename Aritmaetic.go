Algorithm:
Step 1: Start.
Step 2: Declare variables num1, num2 and sum.
Step 3: Read values for num1, num2.
Step 4: Add num1 and num2 and assign the result to a variable sum.
Step 5: Display sum.
Step 6: Stop.




Code:

package main

import "fmt"

func main() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("The Arithmatic operation are:\n ")

	fmt.Printf("The Addition of two numbers is %d\n ", c)
	c = a - b
	fmt.Printf("The Substraction of two numbers is %d\n", c)
	c = a * b
	fmt.Printf("The Multiplication of two numbers is %d\n", c)
	c = a / b
	fmt.Printf("The Division of two numbers is %d\n", c)
	c = a % b
	fmt.Printf("The Modules of two numbers is %d\n", c)

}
