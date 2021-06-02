## Zero value  
1. Declare the variable with `var` keyword but not assign value to it then assign default value of that data type of variable.  

## default with their data type  
1. false for booleans
1. 0 for integers
1. 0.0 for floats
1. "" for strings
1. nil for
    1.  pointers
    1.  functions
    1.  interfaces
    1.  slices
    1.  channels
    1.   maps 
   
```go
package main

import (
	"fmt"
)

// DECLARE a variable to be of a certain TYPE
var name string
var number int
var fnumber float64
var condition bool

func main() {

	fmt.Print("var name string is :\t", name)
	fmt.Printf("\t%T\t\n", name)

	fmt.Print("var number int :\t", number)
	fmt.Printf("\t%T\t\n", number)

	fmt.Print("var fnumber float64  :\t", fnumber)
	fmt.Printf("\t%T\t\n", fnumber)

	fmt.Print("var condition bool :\t", condition)
	fmt.Printf("\t%T\t\n", condition)

}

```  
1.output of program  
```bash
var name string is :            string  
var number int :        0       int     
var fnumber float64  :  0       float64 
var condition bool :    false   bool    
```