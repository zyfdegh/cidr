# CIDR
A package which can calculates the range of IP addresses from a CIDR string.

# Example

192.168.1.0/24 -> 192.168.1.1 ~ 192.168.1.254

```go
package main

import (
	"fmt"
	"log"

	"github.com/zyfdegh/cidr"
)

func main() {
	from, to, err := cidr.GetRange("192.168.1.0/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("from: %s\n", from)
	fmt.Printf("to: %s\n", to)
}

```