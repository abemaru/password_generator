package main

import (
	"fmt"
	"math/big"
	"flag"
	"crypto/rand"
)

func main() {
	p := generatePassword()
	fmt.Printf(p)
}

const rs2Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	lowerCase = "abcdefghijklmnopqrstuvwxyz"

	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	digit = "1234567890"
)

type ParserArguments struct {
	length int
}

func getArguments() *ParserArguments {
	var l = flag.Int("l", 6, "password length")
	flag.Parse()

	g := &ParserArguments{
		length: *l,
	}

	return g 
}

func generatePassword() string {
	p := getArguments()

	b := make([]byte, p.length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(rs2Letters))))
		if err != nil {
			panic(err)
		}
		b[i] = rs2Letters[n.Int64()]
	}
	return string(b)
}
