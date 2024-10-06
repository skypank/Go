package main

import (
	"fmt"
	"reflect"
)

// variable declarations may be "factored" into blocks, like import statements.
// go basic variable types

var (
	b    bool
	s    string
	i    int
	i8   int8
	i16  int16
	i32  int32
	i64  int64
	u    uint
	u8   uint8
	u16  uint16
	u32  uint32
	u64  uint64
	uptr uintptr
	by   byte // alias for uint8
	r    rune // alias for int32, represents a Unicode code point
	f    float32
	f64  float64
	c    complex64 // used for complex number like 1+2i
	c128 complex128
)

// Variables declared without an explicit initial value are given their zero value.
// 0 for int, false for bool, "" (empty string) for strings

func main() {
	// size in bytes, bool is 1 byte
	fmt.Printf("Type: %T Size: %v\n", b, reflect.TypeOf(b).Size())
	// string is 16 bytes
	fmt.Printf("Type: %T Size: %v\n", s, reflect.TypeOf(s).Size())
	// int is 8 bytes
	fmt.Printf("Type: %T Size: %v\n", i, reflect.TypeOf(i).Size())
	fmt.Printf("Type: %T Size: %v\n", i8, reflect.TypeOf(i8).Size())
	fmt.Printf("Type: %T Size: %v\n", i16, reflect.TypeOf(i16).Size())
	fmt.Printf("Type: %T Size: %v\n", i32, reflect.TypeOf(i32).Size())
	fmt.Printf("Type: %T Size: %v\n", i64, reflect.TypeOf(i64).Size())
	// uint is 8 bytes
	fmt.Printf("Type: %T Size: %v\n", u, reflect.TypeOf(u).Size())
	fmt.Printf("Type: %T Size: %v\n", u64, reflect.TypeOf(u64).Size())
	fmt.Printf("Type: %T Size: %v\n", uptr, reflect.TypeOf(uptr).Size())
	//The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	fmt.Printf("Type: %T Size: %v\n", f, reflect.TypeOf(f).Size())
	fmt.Printf("Type: %T Size: %v\n", c128, reflect.TypeOf(c128).Size())

	// assignment between items of different type requires an explicit conversion.
	f64 = float64(i)
	// convert i of type int to f of type float64, if we remove the explicit float64, this will error out

	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	// When the right hand side of the declaration is typed, the new variable is of that same type
	var i int
	k := i // k is an int
	fmt.Printf("k type %T\n", k)
	// But when the right hand side contains an untyped numeric constant,
	j := 1.1
	fmt.Printf("j type %T\n", j) // j is float64
	// the new variable type will depend on the precision of the constant

	// Constants are declared like variables, but with the const keyword.
	// Constants can be character, string, boolean, or numeric values.
	const T = true
	fmt.Printf("value of t %v", T)
	//Constants cannot be declared using the := syntax. because := is meant for variable and const are immutable
}
