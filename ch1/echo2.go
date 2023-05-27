// Echo2 prints its command-line arguments.
package main
import (
"fmt"
"os"
"time"
)

func main() {
start := time.Now()
s, sep := "", ""
for _, arg := range os.Args[0:] {
s += sep + arg
sep = " "
}

for _, arg := range os.Args[0:]{

fmt.Println(arg)
}
secs := time.Since(start).Seconds()

fmt.Printf("%.9fs", secs)
}