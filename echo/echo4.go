// Echo function that prints all os.Args inputs
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.8fs elaspsed \n", time.Since(start).Seconds())
}
