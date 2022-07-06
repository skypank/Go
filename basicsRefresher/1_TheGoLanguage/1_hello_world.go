package main
/* package keyword signifies that main is collection of source code file
 You can add mulitple files with this package name and get a single binary output 
 Each file in package should contain first line as name of package

 There are two types of packages in go
 1. Executable : These packages generate file that we can run
 2. Reusable : These packages are used as helper code, doesn't generate any binary

 The name of the package determines if the package is executable  or reusable.
 All executable package have name 'main'
 All other packages with any other name means it is a reusable package with helper code and wont' generate executables

 Also each package 'main' MUST have a function named 'main' in it */

import (
  "fmt"
)
/* import keyword provides access of underlying package, in this case "fmt", to the current package 'main'.
 'fmt' is a standard package, similarly user created packages can be imported by 'import' command. 
 
 List of all standard packages available in go is listed here - https://pkg.go.dev/std */

func main() {
  fmt.Println("Hello World")
  // Println is a function available under package fmt and is available for export
  // Notice that, the name of fucntion begins with capital letter 'Println'
  // In go, a function from a package can be exported only if it defined with capital letter
  // if package define func with small letter, then such function cannot be imported by other packages.
  // This is go's way of defining accessibility similar to private and public function in other languages.
}
/* 'func' keyword is short form of function
 syntax :
         func nameOfFunc (listOfArgs) returnType {
             code...
             return valueType
             } 
             */

/* Go file organisation :
   package declaration -> package 'main'
   import packages     -> import "fmt"
   declare func & code -> func main ...  */
