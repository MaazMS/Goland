## expression  
1.  An expression is a code that computes one or multiple values.  
1. Unlike statements an expression can not stand on its own on a code line. it should be used with or within a statement   
   or with an operator.  
1. It is because an expression returns a value.    
1. here what are the expression  
```go
package main

import "fmt" 

func main()  {
    fmt.Println("Hello")
}
``` 
1. hello is an expression because it returns.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
``` 