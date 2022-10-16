package main

import (
	"fmt"
	"flag"
)

func main() {
	generator()
}

func generator() {
	var (
		l = flag.Int("l", 6, "password length")
	)
	flag.Parse()
	fmt.Printf("param -l : %s\n", *l)
}
