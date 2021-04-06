## 
1. godoc no longer works, so you have two options now.  
1. option one use `go doc` command instead.    
1. option two colorful output, use Inanc Gumus's custom package: `godocc`.  

1. It creates documentation from code comments automatically.  
1. When you start a comment on just above a code line that has a package clause or a function declaration variable  
declaration and so on. Go will use your comment to create the documentation automatically.   
1. 

## Example   
1. Program `countCPU.go` to count number of CPU in a machine .  
1. `go doc -src runtime NumCPU`  
1. go doc -src package_name function_name     
1. `-src` means show me the source code.  
