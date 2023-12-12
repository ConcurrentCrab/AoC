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

var cubeNums map[string]int

func main() {
	s := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		cubeNums = make(map[string]int)
		_, a, _ := strings.Cut(str, ": ")
		for _, pull := range strings.Split(a, "; ") {
			for _, pullc := range strings.Split(pull, ", ") {
				n, c, _ := strings.Cut(pullc, " ")
				cubeNums[c] = max(cubeNums[c], atoi(n))
			}
		}
		s += cubeNums["red"] * cubeNums["green"] * cubeNums["blue"]
	}
	fmt.Println(s)
}
