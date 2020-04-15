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

	str := strings.Join(ans, "\n")
	fmt.Println(str)
}

var ans = make([]string, 0)

func isPrime(n int) {
	if n >= 1000000 {
		for _, p := range primes {
			if p >= 46340 || p*p > n {
				ans = append(ans, fmt.Sprintf("%d", 1))
				//println(1)
				return
			} else if n%p == 0 {
				ans = append(ans, fmt.Sprintf("%d", 0))
				//println(0)
				return
			}
		}
	} else {
		if mark[n] {
			ans = append(ans, fmt.Sprintf("%d", 0))
			//println(0)
		} else {
			ans = append(ans, fmt.Sprintf("%d", 1))
			//print(1)
		}
	}
}
