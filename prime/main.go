package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var mark = make([]bool, 1000000)
var primes = make([]int, 78498)

func main() {
	path := os.Args[1]
	mark[0] = true
	mark[1] = true
	p := 0
	for i := 2; i < 1000000; i++ {
		if mark[i] {
			continue
		}
		for j := i * 2; j < 1000000; j += i {
			mark[j] = true
		}
		primes[p] = i
		p++
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		isPrime(n)
	}
}

func isPrime(n int) {
	if n >= 1000000 {
		for _, p := range primes {
			if p > 47000 {
				fmt.Println(1)
				return
			} else if n%p == 0 {
				fmt.Println(0)
				return
			}
		}
	} else {
		if mark[n] {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
