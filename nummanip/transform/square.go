package transform

//return square of slice
func SquareSlice(s []int) []int {
	
	//make slice of length len(s)
	squareS := make([]int, len(s))

	for index, value := range s {
		squareS[index] = value * value
		
	}

	//return transformed slice
	return squareS
}