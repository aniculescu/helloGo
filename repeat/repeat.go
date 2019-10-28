package iteration

import "fmt"

const repeatCount = 5

func Repeat(repeatString string, repeatTimes int) string {
	var result string
	for i := 0; i < repeatTimes; i++ {
		result += repeatString
	}
	return result
}

func main() {
	fmt.Println(Repeat("a", 2))
}
