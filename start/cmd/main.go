package main

import (
	"fmt"
	"os"

	"hello"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
