package main

import (
	"fmt"
	"strings"
)

// A string is a slice of bytes in Go.
// Strings can be created by enclosing a set of characters inside double quotes " ".

func main() {
	// declare and assign using :=
	variable := "this is string"
	fmt.Printf("string value : %s", variable)
	// Since string are slice of bytes, you can access individual bytes of string using a loop construct
	// this for loop start with first byte of string and goes till len of string by stepping up one
	for i := 0; i < len(variable); i++ {
		// %v is used to print the value of each byte in ascii
		fmt.Println()
		fmt.Println("ascii, hex, char")
		fmt.Printf("%v    ", variable[i])
		// %x is used to print the value of each byte in hex
		fmt.Printf("%x    ", variable[i])
		// %c is used to print the value of each byte in char
		fmt.Printf("%c    ", variable[i])
		fmt.Println()
	}

	// Strings in Go are Unicode compliant and are UTF-8 Encoded.
	// Above code we are assuming that each char is one byte, and accesing it byte by byte
	// Since many unicode character can be longer than one byte, and go is unicode compliant,
	// we have to find a way for accessing character longer than one byte

	variable = "Señor"
	for i := 0; i < len(variable); i++ {
		fmt.Printf("%c    ", variable[i])
	}

	//  the output of this segment will look like
	//  S    e    Ã    ±    o    r
	// you can see that since one byte is not enough to hold character like ñ
	// we are having wrong output, to overcome this go has another specifier where each character can hold more than one byte
	// we call it rune, each element in a rune can hold upto int32 size
	// converting the above variable to a slice of runes
	runes := []rune(variable)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	// now the output is
	// S e ñ o r
	// as expected

	// we have a better version of foor loop in go, where looping over is simplified
	for i, rune := range variable {
		fmt.Printf("\nindex is %d and char is %c\n", i, rune)
	}
	// the outpute will look like
	// index is 0 and char is S
	// index is 1 and char is e
	// index is 2 and char is ñ
	// index is 4 and char is o
	// index is 5 and char is r
	// you can see that the special character has take two bytes and next character has occupied later bytes

	// converting a byte into string
	byteSlice := []byte{116, 104, 105, 115, 32, 105, 115, 32, 115, 116, 114, 105, 110, 103}
	// string function is used for conversting bytes into slice
	str := string(byteSlice)
	fmt.Println(str)
	// similarly for rune slice, same string func can be used for conversion

	// string functions
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	// string comparison
	string1 := "Hello"
	string2 := "world"
	// '==' is comparision operator
	if string1 == string2 {
		fmt.Print("strings are equal")
	}
	// string concatenation
	// '+' is the concatenation operator
	result := string1 + " " + string2
	fmt.Println(result)
	// alternatively, we can use fmt's sprintf function
	result = fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)

	// Strings are immutable in Go. Once a string is created it's not possible to change it.
	// h := "hello"
	// trying to change h[0]='t' will result in error cannot assign to h[0]

}
