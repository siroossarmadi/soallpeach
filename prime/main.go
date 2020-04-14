package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(num int) int {
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
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%v\n", isPrime(num))
	}
}
