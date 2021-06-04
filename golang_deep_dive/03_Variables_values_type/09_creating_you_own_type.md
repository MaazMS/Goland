## Creating  your own type  
1. golang is static type language,    
1. Their for it is type is more important in golang.   
1. Create our own type with the help of `type` keyword.  
1. you mast be defined underline type of your type.   

## Types
1. https://golang.org/ref/spec#Types  
1. Each type T has an underlying type: If T is one of the predeclared boolean, numeric, or string types, or a type literal,    
   the corresponding underlying type is T itself.  
1.  Otherwise, T's underlying type is the underlying type of the type to which T refers in its type declaration.  

```go
type T1 string
type T2  T1
```
1. The underlying type of string T1, T2 is string.  

```go
package main

import (
	"fmt"
)

var a int

// create our own type palindrome and underline type is int
type palindrome int

// create variable b with TYPE palindrome
var b palindrome

func main() {
	a = 15
	// assign value to variable b with TYPE palindrome is int 25
	b = 25
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}

```