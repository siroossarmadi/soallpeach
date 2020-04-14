package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	var perline int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		if big.NewInt(int64(perline)).ProbablyPrime(0) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
