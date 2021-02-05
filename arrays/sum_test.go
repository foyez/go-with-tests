package arrays

import (
	"reflect"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	// assertCorrectMessage := func(t *testing.TB, got, want int) {
	// 	t.Helper()

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// }
	
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// numbers := [5]int{1, 2, 3, 4, 5} // numbers := [...]int{1, 2, 3, 4, 5} - array
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}
