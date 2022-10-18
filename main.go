package main

import (
	"fmt"
	"math/big"
	"flag"
	"crypto/rand"
)

func main() {
	p := generator()
	fmt.Printf(p)
}

const rs2Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generator() string {
	var pwLen = flag.Int("l", 6, "password length")
	
	flag.Parse()

	b := make([]byte, *pwLen)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(rs2Letters))))
		if err != nil {
			panic(err)
		}
		b[i] = rs2Letters[n.Int64()]
	}
	return string(b)
}
