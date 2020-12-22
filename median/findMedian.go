package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Definition:

You are given a feature to display the median price of a given beverage type. Given an n sized unsorted array of integer prices, write a function findMedian([]) to return the median price. You can assume that the array will contain at least one element.

*/

//FindMedian function
func FindMedian(prices []int) (median float64) {
	var middle int = 0

	//for arrays of 1 element, return
	if len(prices) < 2 {
		return float64(prices[0])
	}
	//calculate the key of the middle element in the array
	middle = int(math.Ceil(float64(len(prices)-1) / 2))

	//sort array
	sort.Ints(prices[:])

	//check if odd or even number of elements
	if len(prices)%2 > 0 {
		//array with an odd number of elements
		//the median is the middle number
		median = float64(prices[middle])
	} else {
		//array with an even number of elements
		//the median is the average of the middle two numbers
		num1 := prices[middle]
		num2 := prices[middle-1]

		median = float64((num1 + num2)) / 2
	}

	return median
}

func main() {
	testArray1 := []int{1, 6, 3, 5, 8, 9, 4, 10, 2}
	testArray2 := []int{1, 6, 3, 5, 8, 9, 4, 10, 2, 7}
	fmt.Println("Prices List: ", testArray1)
	fmt.Println("Median Price: ", FindMedian(testArray1))
	fmt.Println("Prices List: ", testArray2)
	fmt.Println("Median Price: ", FindMedian(testArray2))
}
