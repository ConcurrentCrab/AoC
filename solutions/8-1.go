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

func main() {
	network := make(map[string]([2]string))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p := scanner.Text()
	path := slicesMap([]byte(p), func(b byte) int {
		switch b {
		case 'L':
			return 0
		case 'R':
			return 1
		}
		return -1
	})
	scanner.Scan()
	for scanner.Scan() {
		str := scanner.Text()
		b, a, _ := strings.Cut(str, " = ")
		a = strings.Trim(a, "()")
		network[b] = [2]string(strings.Split(a, ", "))
	}
	s := 0
	cur, i := "AAA", 0
	for true {
		dir := path[i]
		cur = network[cur][dir]
		i = (i + 1) % len(path)
		s++
		if cur == "ZZZ" {
			break
		}
	}
	fmt.Println(s)
}
