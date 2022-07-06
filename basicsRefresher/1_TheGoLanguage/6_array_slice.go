package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")
	var a [5]int
	// declares a variable a as an array of 5 integers.
	fmt.Println(a)
	// An array's length is part of its type, so arrays cannot be resized.
	// All elements in an array are automatically assigned the zero value of the array type
	// The index of an array starts from 0 and ends at length - 1.
	a[0] = 1 // starts
	a[4] = 4 // ends (length - 1)
	fmt.Println(a)
	// shorthand declaration
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	// few elements assignment
	c := [5]int{2, 3}
	// only first two elements will be assigned, rest will hv zero value
	fmt.Println(c)
	// no explicit size needed, it will be automatically calculated
	d := [...]int{1, 2, 3}
	fmt.Println(d)
	// copy among diff array lenth not allowed, even if they are of same variable
	// c = d
	// will result in error as c is of type [5]int and d is of type [3]int
	// The size of the array is a part of the type
	c = b
	// this will work because both c and b are of type [5]int
	fmt.Println(c)
	// Arrays in Go are value types and not reference types.
	// This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable.
	// now changes in b will not effect c
	b[0] = 0
	fmt.Println(c)
	fmt.Println(b)
	// Even when array is passed to func, its pass by value, so values remains unchanged

	// Multidimensinal array
	var multi [2][2]int
	fmt.Println(multi)
	// declare and initialize
	multiA := [2][2]int{{2, 3}, {4, 5}}
	fmt.Println(multiA)
	multiB := [3][2]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	fmt.Println(multiB)

	// Slices -  A slice is a dynamically-sized, flexible view into the elements of an array.
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[0:3] // this gives access of element from 0th to 2nd index, last element is excluded.
	fmt.Println(slice)
	// slice can be directly created using []T syntax
	sliceArr := []int{1, 2, 3} // this creates array of three values and returns a slice
	// the difference between this and array declaration is size specifier inside []
	// [3]int{1,2,3} - this will create and return array of size 3
	// []int{1,2,3} - this will create array but return slice on array of size 3
	fmt.Println(sliceArr)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	nameArr := [4]string{"amar", "akbar", "anthony", "gurdeep"}
	slice1 := nameArr[0:2]
	slice2 := nameArr[2:4]
	slice1[0] = "shreyas" // this will modify not only this slice1, but the nameArr 1st element as well,
	// since the slice is just reference to underlying array
	slice2[0] = "tony"
	fmt.Println(slice1, slice2, nameArr)
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	slice3 := nameArr[1:]
	slice4 := nameArr[:4]
	fmt.Println(slice3, slice4)
	// The default is zero for the low bound and the length of the slice for the high bound.

	// length and capacity
	fmt.Printf("length of slice: %d, capacity of slice: %d", len(slice1), cap(slice1))
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	fmt.Printf("length of slice: %d, capacity of slice: %d", len(slice2), cap(slice2))
	// the cap of slice1 and slice2 differes here since slice1 begind at first element, and slice2 begins at 3rd
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

	// nil slice
	// The zero value of a slice is nil.
	nilSlice := []int{}
	fmt.Println("\nnil value of slice", nilSlice)
	// A nil slice has a length and capacity of 0 and has no underlying array.
}
