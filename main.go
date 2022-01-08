package main

import (
	"fmt"

	"github.com/leson/go_poc/greetings"
)

func main() {
	fmt.Println("invoke self-define module of go!")
	fmt.Println(greetings.Hello("Leson"))
}
