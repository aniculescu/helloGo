package integers

//Adder function
//Takes two int inputs: n1, n2
//Returns: n1 + n2
func Adder(n1, n2 int) int {
	return n1 + n2
}

//Sum Function
func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

//SumAll Function
//Takes 1 or more slices
//Returns: a slice
func SumAll(slices ...[]int) []int {
	var totals []int
	for _, slice := range slices {
		totals = append(totals, Sum(slice))
	}
	return totals
}

//SumAllTrails Function
func SumAllTrails(slices ...[]int) []int {
	var totals []int
	//var tail []int
	mySum := 0
	for _, slice := range slices {
		if len(slice) > 0 {
			mySum = Sum(slice[1:])
		} else {
			mySum = 0
		}
		totals = append(totals, mySum)
	}
	return totals
}
