package main

import "fmt"

func main() {
	numOfBirds := 20
	ptrNumOfBirds := &numOfBirds

	fmt.Println(*ptrNumOfBirds) // 20

	numOfBirds = 10 // change the value of the variable

	fmt.Println(*ptrNumOfBirds)
}
