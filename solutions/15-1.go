//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func slicesReduce[T, U any](ts []T, u U, f func(T, U) U) U {
	for i := range ts {
		u = f(ts[i], u)
	}
	return u
}

func hashStr(s string) int {
	b := []byte(s) // iterating over a string returns runes not bytes
	n := 0
	for _, v := range b {
		n += int(v)
		n *= 17
		n &= (1 << 8) - 1 // mod 256
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	vals := slicesMap(strings.Split(str, ","), hashStr)
	sum := slicesReduce(vals, 0, func(a, b int) int { return a + b })
	fmt.Println(sum)
}
