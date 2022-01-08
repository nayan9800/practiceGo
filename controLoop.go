package main

import (
	"fmt"
	"math"
)

func controlAndLoops() {

	/*defer  is used to execute something in last*/
	defer fmt.Println("Bye Bye never see u again 😐")
	fmt.Println("Go loops")
	fmt.Println(sum(10))
	fmt.Println(sum2(1000))

	fmt.Println(math.Log2(25.0))
	/*If-else in Go*/
	no := 4465
	if no%2 == 0 {
		fmt.Printf("No %d is Even\n", no)
	} else {
		fmt.Printf("No %d is Odd\n", no)
	}

	/*If statement with short variable statment
	here the no variable is intialized using sum2
	function and then used in if statment*/
	if no := sum2(10); no > 10 {
		fmt.Println(no)
	}

	/*Switch in Go*/
	/*case does not need to be  with break*/
	no = 1
	fmt.Printf("%d in Romman is ", no)
	switch no {
	case 1:
		fmt.Println("I")
	case 2:
		fmt.Println("II")
	case 3:
		fmt.Println("II")
	default:
		fmt.Println("Sorry but i Don't know")
	}

	/*switch with no condition acts as if-else block*/
	no = -65
	switch {
	case no < 0:
		fmt.Printf("%d is negative\n", no)

	case no > 0:
		fmt.Printf("%d is positive\n", no)

	default:
		fmt.Printf("%d is zero\n", no)
	}
}

/*for loop in Go has
1. a intial variable state for example i:= 0
2. a condition expression something like i<10
3. and a iteraive statment or post statment*/
func sum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum
}

/*intial and post statement is optional
  and use is at as while loop in C
*/
func sum2(n int) int {
	sum := 1

	for sum < n {
		sum += sum
	}
	/*without conditional statement it will become
	infinte loop
		for {

		}
	*/
	return sum
}
