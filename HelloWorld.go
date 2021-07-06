// to compile: go build <file>
// to run: ./<executable>
// to compile and run (say after editing your file): go run <file> (this does not make an executable file)

// package declaration: part of every program, informs compiler whether to create an executable or library
// package main means that the file can create an executable
package main
// other packages: https://golang.org/pkg/

import "fmt"
// to compile: go build <file>
// to run: ./<executable>
// to compile and run (say after editing your file): go run <file> (this does not make an executable file)

// package declaration: part of every program, informs compiler whether to create an executable or library
// package main means that the file can create an executable
package main
// other packages: https://golang.org/pkg/

// fumpt package: for printing and formatting data
import "fmt"
import t "time"  // t is our alias
/*
another way to import these:
import(
	import "fmt"
	import t "time"
)
error will be thrown if package is imported but not used (similar to Stencil)
*/

// this function will automatically run
func main() {
	// yaassss no semicolons!!!
	fmt.Println("Hello World")
	// UTC time
	fmt.Println(t.Now())
}

/*
for help:
Go stackoverflow: https://stackoverflow.com/questions/tagged/go
Go docs: https://golang.org/
When Googling, it helps to use "golang" instead of "go" sometimes
Go also has built-in documentation system, heres the commands:
- for package info: go doc <package name>
- for function info: go doc <package name>.<func name>
*/