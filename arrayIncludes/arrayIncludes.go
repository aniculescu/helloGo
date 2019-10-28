package algos

//ArrayIncludes function
func ArrayIncludes(arr1, arr2 []int) []int {
	bigarray := arr1
	smallarray := arr2
	var found []int
	if len(arr1) > len(arr2) {
		bigarray = arr1
		smallarray = arr2
	} else if len(arr1) < len(arr2) {
		bigarray = arr2
		smallarray = arr1
	}
	for _, v := range bigarray {
		if indexOf(v, smallarray) < 0 {
			found = append(found, v)
		}
	}
	return found
}

func indexOf(element int, array []int) int {
	for i, v := range array {
		if element == v {
			return i
		}
	}
	return -1
}
