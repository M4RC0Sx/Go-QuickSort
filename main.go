package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func partition(array []int, indexMin int, indexMax int) int {

	pivot := array[indexMax]
	i := indexMin - 1

	swapper := reflect.Swapper(array)

	for j := indexMin; j <= indexMax-1; j++ {

		if array[j] < pivot {

			i++
			swapper(i, j)
		}
	}

	swapper(i+1, indexMax)

	return i + 1

}

func quickSort(array []int, indexMin int, indexMax int) {

	if indexMin < indexMax {

		m := partition(array, indexMin, indexMax)

		quickSort(array, indexMin, m-1)
		quickSort(array, m+1, indexMax)
	}

}

func main() {

	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var arrayLen int = len(array)

	fmt.Printf("-> Given array: %v\n", array)
	fmt.Printf("-> Given array length: %d\n\n", arrayLen)

	fmt.Println("-> Shuffling array...")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(arrayLen, func(i, j int) { array[i], array[j] = array[j], array[i] })
	fmt.Printf("-> Shuffled array: %v\n\n", array)

	fmt.Println("-> Starting QuickSort algorithm...")
	quickSort(array[:], 0, arrayLen-1)
	fmt.Println("-> QuickSort algorithm finished!")

	fmt.Printf("\n-> Shorted array: %v\n", array)
}
