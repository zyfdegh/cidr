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
