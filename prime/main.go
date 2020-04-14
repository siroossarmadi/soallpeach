package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func isPrime(num int) int {
	if num == 1 {
		return 0
	}
	if num <= 3 {
		return 1
	}
	if num%2 == 0 || num%3 == 0 {
		return 0
	}
	temp := int(math.Sqrt(float64(num))) + 1
	for i := 5; i < temp; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	fileName := os.Args[1]
	file, _ := os.Open(fileName)
	defer file.Close()
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
		fmt.Println(isPrime(perline))
	}
}
