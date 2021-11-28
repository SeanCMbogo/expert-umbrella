// A streamlined approach to the echo function
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Raw output with brackets
// func main() {
// 	fmt.Println(os.Args[1:])
// }
