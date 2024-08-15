// main.go
package main

import (
	"fmt"

	"github.com/mohammedaouamri5/vector/vector"
)





func main() {

	
	bruh := vector.New(10, 0.5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Printf("%+v\n", bruh)  // Uses the Format method with %+v
	fmt.Printf("%v\n")   // Uses the Format method with %v
	fmt.Println(bruh.String()) // Uses the String method
	bruh.Push(11)
	fmt.Printf("%+v\n", bruh)  // After pushing a new element

} 