## what is compiled time?   
1. A program is convert source code (go program ) to binary language is called compiled time.   
## How to create compiled time code?  
1. Write go program.   
```
package main

import "fmt"

func main(){
	fmt.Println("Maaz Shaikh")
}

``` 
1. open terminal enter `go build` 
1. `build       compile packages and dependencies`  

## How to execute compile code?  
./executeble_file_name  

## How to identify an executable file?   
1. ls -l 
`-rwxrwxr-x 1 maaz maaz 2034781 Feb 15 10:50 01-print-names`  -x means it is an executable file.      
2. file file_name   
`file 01-print-names `  

## What is runtime?  
1. To execute compile code is start running.  

## How to perform runtime?  
1. go run file_name.   
` go run main.go`   
1. It is compiled and run program   

## what is behind go run?   
1. go run -x main.go   
2. it is not use executable file.   
step 1: compiled source code.  
step 2: save compiled source code into operating system temporary directory.   
step 3: linking packages to compiled code.  
```
maaz@maaz-Lenovo-G50-70:~/github/LearningGO/code/src/01-write-your-first-program/exercises/01-print-names$ go run -x main.go 
WORK=/tmp/go-build501654689
mkdir -p $WORK/b001/
cat >$WORK/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/home/maaz/.cache/go-build/97/979116683edb683f474d4b0a499295cd503d847303e67d68d065996e5977e2e2-d
packagefile fmt=/usr/local/go/pkg/linux_amd64/fmt.a
packagefile runtime=/usr/local/go/pkg/linux_amd64/runtime.a
packagefile errors=/usr/local/go/pkg/linux_amd64/errors.a
packagefile internal/fmtsort=/usr/local/go/pkg/linux_amd64/internal/fmtsort.a
packagefile io=/usr/local/go/pkg/linux_amd64/io.a
packagefile math=/usr/local/go/pkg/linux_amd64/math.a
packagefile os=/usr/local/go/pkg/linux_amd64/os.a
packagefile reflect=/usr/local/go/pkg/linux_amd64/reflect.a
packagefile strconv=/usr/local/go/pkg/linux_amd64/strconv.a
packagefile sync=/usr/local/go/pkg/linux_amd64/sync.a
packagefile unicode/utf8=/usr/local/go/pkg/linux_amd64/unicode/utf8.a
packagefile internal/bytealg=/usr/local/go/pkg/linux_amd64/internal/bytealg.a
packagefile internal/cpu=/usr/local/go/pkg/linux_amd64/internal/cpu.a
packagefile runtime/internal/atomic=/usr/local/go/pkg/linux_amd64/runtime/internal/atomic.a
packagefile runtime/internal/math=/usr/local/go/pkg/linux_amd64/runtime/internal/math.a
packagefile runtime/internal/sys=/usr/local/go/pkg/linux_amd64/runtime/internal/sys.a
packagefile internal/reflectlite=/usr/local/go/pkg/linux_amd64/internal/reflectlite.a
packagefile sort=/usr/local/go/pkg/linux_amd64/sort.a
packagefile math/bits=/usr/local/go/pkg/linux_amd64/math/bits.a
packagefile internal/oserror=/usr/local/go/pkg/linux_amd64/internal/oserror.a
packagefile internal/poll=/usr/local/go/pkg/linux_amd64/internal/poll.a
packagefile internal/syscall/execenv=/usr/local/go/pkg/linux_amd64/internal/syscall/execenv.a
packagefile internal/syscall/unix=/usr/local/go/pkg/linux_amd64/internal/syscall/unix.a
packagefile internal/testlog=/usr/local/go/pkg/linux_amd64/internal/testlog.a
packagefile sync/atomic=/usr/local/go/pkg/linux_amd64/sync/atomic.a
packagefile syscall=/usr/local/go/pkg/linux_amd64/syscall.a
packagefile time=/usr/local/go/pkg/linux_amd64/time.a
packagefile internal/unsafeheader=/usr/local/go/pkg/linux_amd64/internal/unsafeheader.a
packagefile unicode=/usr/local/go/pkg/linux_amd64/unicode.a
packagefile internal/race=/usr/local/go/pkg/linux_amd64/internal/race.a
EOF
mkdir -p $WORK/b001/exe/
cd .
/usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/main -importcfg $WORK/b001/importcfg.link -s -w -buildmode=exe -buildid=253_2vCBbo2CzYHT-w8d/As4iW_iUknNWbuzcB-4Z/cWKCQoB0HEL_SU527Shd/253_2vCBbo2CzYHT-w8d -extld=gcc /home/maaz/.cache/go-build/97/979116683edb683f474d4b0a499295cd503d847303e67d68d065996e5977e2e2-d
$WORK/b001/exe/main
Maaz Shaikh

```

## What is the difference between `go build` and `go run`?   
2. `go run` both compiles and runs a program; whereas `go build` just compiles it.   
## Go saves the compiled code in a directory. What is the name of that directory?
1. The same directory where you call `go build`   
