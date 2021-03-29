package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Find all positive integer pairs a,b such that a + b = M, a xor b = N\n")
		fmt.Fprintf(os.Stderr, "usage: %s M N\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "M, N hexadecimal representation\n")
		return
	}
	M, err := strconv.ParseUint(os.Args[1], 0x10, 64)
	if err != nil {
		log.Fatal(err)
	}
	N, err := strconv.ParseUint(os.Args[2], 0x10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var a uint64
	for a = 0; a <= M/2; a++ {
		b := M - a
		if a^b == N {
			fmt.Printf("%x\t%x\t%x\t%x\n", a, b, a+b, a^b)
		}
	}
}
