//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	s := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		a, b := 0, 0
		for i := 0; i < len(str); i++ {
			if isDigit(str[i]) {
				a = int(str[i] - '0')
				break
			}
		}
		for i := len(str) - 1; i > -1; i-- {
			if isDigit(str[i]) {
				b = int(str[i] - '0')
				break
			}
		}
		s += a * 10
		s += b
	}
	fmt.Println(s)
}
