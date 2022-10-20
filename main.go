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

const (
	lowerCase = "abcdefghijklmnopqrstuvwxyz"

	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	digit = "1234567890"

	symbol = "!@#$"
)

type ParserArguments struct {
	length int
	upperCase bool
	numbers bool
	symbols bool
}


func getArguments() *ParserArguments {
	var l = flag.Int("l", 6, "password length")
	var u = flag.Bool("u", true, "generate password with upper cases")
	var n = flag.Bool("n", true, "generate password with numbers")
	var s = flag.Bool("s", true, "generate password with symbols")
	flag.Parse()

	g := &ParserArguments{
		length: *l,
		upperCase: *u,
		numbers: *n,
		symbols: *s,
	}

	return g 
}


func  generateStrings(isUpperCase bool, isNumbers bool, isSymbols bool) string {
	letters := lowerCase
	if isUpperCase {
		letters += upperCase
	}

	if isNumbers {
		letters += digit
	}

	if isSymbols {
		letters += symbol
	}

	return letters
}

func generatePassword() string {
	p := getArguments()
	l := generateStrings(p.upperCase, p.numbers, p.symbols)


	b := make([]byte, p.length)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(l))))
		if err != nil {
			panic(err)
		}
		b[i] = l[n.Int64()]
	}
	return string(b)
}
