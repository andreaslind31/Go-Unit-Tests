package transform

import (
	"reflect"
	"testing"
)

//test case for Square func in square.go
func TestTransformSquare( t *testing.T)  {
	testSlice := []int{1,2,3,4,5,6,7,8,9}
	expectedResult := []int{1,4,9,16,25,36,49,64,81}

	result := SquareSlice( testSlice)

	if reflect.DeepEqual(expectedResult, result) {
		t.Log("SquareSlice PASSED")
	} else {
		t.Errorf("SquareSlice FAILED, expected %v but got %v", expectedResult, result)
	}
}