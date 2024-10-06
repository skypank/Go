package main

import "fmt"

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(newadd(42, 13))
	noArgument()
	fmt.Println(oneArgument())
	fmt.Println(moreArgument())
	fmt.Println(namedArgument())
	insideFunc()
}

// A function can take zero or more arguments.
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last
// here arg type for x is ommited
func newadd(x, y int) int {
	return x + y
}

// a func can return zero or more arguments
func noArgument() {
	fmt.Println("No argument")
}

// return one argument
func oneArgument() string {
	return "one Argument"
}

// more argument
func moreArgument() (string, string, string) {
	return "one", "two", "three"
}

// Named return/ Naked return
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// this should be used only in short functions, otherwise it can be confusing
func namedArgument() (x, y int) {
	x = 1
	y = 2
	return
}

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
var x = 2

func insideFunc() {
	// accessing global var x
	fmt.Println(x)
	// declaring local var x, now the global var x will not be available inside this func
	x := 1
	fmt.Println(x)
}
