//echo's from command-line
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
fmt.Println(s)
secs := time.Since(start).Seconds()
fmt.Printf("time is %.9f", secs)
}