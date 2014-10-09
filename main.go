package main

import (
	"coin"
	"fmt"
)

func main() {
	c := coin.New()
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", c.Flip())
	}
}
