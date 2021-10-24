package main

import (
	"exercism_learn/go/clock"
	"fmt"
)

func main() {
	fmt.Println(clock.New(1, -160))
	fmt.Println(clock.New(1, -40))
	fmt.Println(clock.New(2, 20).Subtract(3000))
}
