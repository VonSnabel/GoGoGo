package main

import "fmt"

type Kek struct {
	x, y int
}

func main() {
	lel := Kek{}
	lel.x = 2

	fmt.Printf("Type: %T. Value: %v", lel, lel)
}
