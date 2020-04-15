package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const chacheSize = 1000000

var mark = make([]bool, 1000000)
var primes = make([]int, 78498)
var ans = make([]string, 0)
var cache = make([]int, chacheSize)

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
		if n > chacheSize {
			ans = append(ans, isPrime(n))
		} else {
			if cache[n] == 0 {
				ip := isPrime(n)
				if ip == "1" {
					ans = append(ans, "1")
					cache[n] = 1
				} else if ip == "0" {
					cache[n] = 2
					ans = append(ans, "0")
				}
			} else if cache[n] == 1 {
				ans = append(ans, "1")
			} else if cache[n] == 2 {
				ans = append(ans, "0")
			}
		}

	}

	str := strings.Join(ans, "\n")
	fmt.Println(str)
}

func isPrime(n int) string {
	if n >= 1000000 {
		for _, p := range primes {
			if p >= 46340 || p*p > n {
				// ans = append(ans, fmt.Sprintf("%d", 1))
				return "1"
			} else if n%p == 0 {
				// ans = append(ans, fmt.Sprintf("%d", 0))
				//println(0)
				return "0"
			}
		}
	} else {
		if mark[n] {
			// ans = append(ans, fmt.Sprintf("%d", 0))
			return "0"
		}
	}
	// ans = append(ans, fmt.Sprintf("%d", 1))
	return "1"

}
