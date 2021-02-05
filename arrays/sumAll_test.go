package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sums)
	// Output: [3 9]
}