package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var mark = make([]bool, 1000000)
var primes = make([]int, 78498)
var lines []string

func sieve(ex chan string) {
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
	ex <- "ex"
}

func read(path string, end chan string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines = strings.Split(string(b), "\n")

	end <- "ex"
}

func main() {
	path := os.Args[1]
	ex := make(chan string, 2)
	go read(path, ex)
	go sieve(ex)
	<-ex
	<-ex
	var ans = make([]byte, len(lines)*2)
	j := 0
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if n >= 1000000 {
			for _, p := range primes {
				if p >= 46340 || p*p > n {
					ans[j] = 49
					ans[j+1] = 10
					break
				} else if n%p == 0 {
					ans = append(ans, 48)
					ans = append(ans, 10)
					break
				}
			}
		} else {
			if mark[n] {
				ans[j] = 48
				ans[j+1] = 10
			} else {
				ans[j] = 49
				ans[j+1] = 10
			}
		}
		j += 2
	}
	os.Stdout.Write(ans)

}
