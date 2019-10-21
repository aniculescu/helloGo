package integers

//Adder function
//Takes two int inputs: n1, n2
//Returns: n1 + n2
func Adder(n1, n2 int) int {
	return n1 + n2
}

func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
