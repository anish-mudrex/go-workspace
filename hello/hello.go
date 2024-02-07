package main

import (
	"fmt"

	"github.com/anish-mudrex/go-workspace/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(12345))
}
