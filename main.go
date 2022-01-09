package main

import (
	"fmt"

	"github.com/leson/go_poc/greetings"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("invoke self-define module of go!")
	fmt.Println(greetings.Hello("Leson"))
	fmt.Println(split(9))
}
