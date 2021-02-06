package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAYS
	//  An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types
	var arr [4]int
	arr[0] = 1
	fmt.Println((arr))

	// compiler count the array elements
	arr1 := [...]int{1, 5}
	fmt.Println(arr1)

	// check types of variable
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(arr).Kind())

	// SLICES
	// []T
	// a slice type has no specified length
	slice := []int{1, 2}

	// check types of variable
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(slice).Kind())

	// create a slice
	// make([]T, len, cap)
	// cap - the maximum length of the segment
	// The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.
	arrCopy := make([]int, len(slice), len(slice) + 1)
	copy(arrCopy, slice)
	arrAppend := append(arrCopy, 3)

	// b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// // b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
	// // b[:2] == []byte{'g', 'o'}
	// // b[2:] == []byte{'l', 'a', 'n', 'g'}
	// // b[:] == b

	fmt.Println(arrCopy)
	fmt.Println(len(arrCopy), cap(arrCopy))
	fmt.Println(arrAppend)
}
