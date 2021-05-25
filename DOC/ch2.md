### Language Syntax
variable 

memory allocation it cos 


How memory is allocate by go 
1. Look for our machine architecture the size of pointer or size of address. 
example if i am use 64bit archtecture then the pointer and address size is 64 bit. 
2. In go have word size which match the address , pointer size of our machine.
3. If address size b4bit, pointer size is 64bit and word size is 64bit the int size is 64bit.  


* concept of ply groud and local machine 
play groud archtecture size is 32bit, then the address size is 32bit, pointer size is 32bit, int size is 32bit. 
* But if run in our machine which is 64bit archtecture then int size is change to 64bit.   

* zero vlue
1. evrey singel variable must be initialize.
2. If we not initialize value ourself then it initialize it's Zero value not default.  
3. example 4byte memory allocation set zero value to every single bit.  
4. one exception for zero value which is string. 
5.`string` is two word archtecure. Two word archtecture has two part. Each word have sperat value. 
6. first word is represent pointer to backing the array of bytes.(which is nil)   
7. second word is reprent length.  
8. **In go consider string zero value be empty.**



var a int 
fmt.print("variable int", a)

short variable declaration. 
a := 10 

casting and conversion 
Go have no casting. 
What is casting with example?.    
I am allocating 10 int for 1byte and I am say to compiler to extends the size of this int 10 for 1byte to 4byte please. Then compiler is increase 10 int 1byte to 4byte.
(exaample of strcut on c means continue memory allocating)   
It create integration issue.   
What is conversion with example?.  
I am allocating 10 int for 1byte and I want to increase the size of int then it create
sperate memory for 4byte.     
aaa := int32(10)