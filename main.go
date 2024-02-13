package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed

	rSlice := make([]int, numItems)

	for i := 0; i < numItems; i++ {

		rSlice[i] = random.Intn(max)

	}

	return rSlice

}

func printSlice(slice []int, numItems int) {

	if len(slice) <= numItems {
		fmt.Println(slice)

	} else {
		fmt.Println(slice[:numItems])
	}

}

func checkSorted(slice []int) {

	for i := 1; i < len(slice); i++ {

		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")

		}
	}
	fmt.Println("The slice is sorted!")

}

func bubbleSort(slice []int) {

	for i := 0; i < len(slice)-1; i++ {
		for j := i; j < len(slice)-i-1; j++ {

			if slice[j] > slice[j+1] {

				slice[j], slice[j+1] = slice[j+1], slice[j]

			}

		}
	}

}

