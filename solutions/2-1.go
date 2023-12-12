//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var cubeNums = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	s := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		f := true
		b, a, _ := strings.Cut(str, ": ")
		for _, pull := range strings.Split(a, "; ") {
			for _, pullc := range strings.Split(pull, ", ") {
				n, c, _ := strings.Cut(pullc, " ")
				if atoi(n) > cubeNums[c] {
					f = false
					goto endline
				}
			}
		}
	endline:
		if f {
			_, id, _ := strings.Cut(b, " ")
			s += atoi(id)
		}
	}
	fmt.Println(s)
}
