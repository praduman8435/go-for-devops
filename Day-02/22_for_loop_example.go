package main

import (
	"fmt"
)

func main() {
	fmt.Println("Basic for Loop:")
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		fmt.Printf("%d ",i)
	}
	fmt.Println("\n")

	fmt.Println("for as a while Loop:")
	i := 0
	for i < 5 {
		// fmt.Println(i)
		fmt.Printf("%d ",i)
		i++
	}
	fmt.Println("\n")



	fmt.Println("Infinite for Loop:")
	j := 0
	for {
		// some code
		if j > 100 {
			break // break out of the loop
		}
		fmt.Printf("%v ",j)
		j++
	}
	fmt.Println("\n")

	fmt.Println("for Loop with range (for slices, arrays, maps, strings):")
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("for Loop with range (for strings):")
	for index, runei := range "hello" {
		fmt.Printf("Index: %d, Rune: %c\n", index, runei)
	}

}
