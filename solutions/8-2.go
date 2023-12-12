//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func slicesFilter[T any](ts []T, f func(T) bool) []T {
	us := make([]T, 0, len(ts))
	for _, v := range ts {
		if f(v) {
			us = append(us, v)
		}
	}
	return us
}

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// homegrow maps.Keys() impl until it's stabilised
func mapsKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}
	return r
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
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
	cur, i := slicesFilter(mapsKeys(network), func(s string) bool { return s[2] == 'A' }), 0
	// finding solutions by travelling all paths together takes impossibly long
	// instead find path length for each individual start and find their LCM
	// this is actually a terrible puzzle because the only viable strategy relies on assumptions on input not explicitly stated:
	//   - the paths all form cycles at the end
	//   - there is exactly one valid endpoint in each cycle
	//   - path length to valid endpoint == cycle length
	//   - path lengths are multiples of input cycle length
	// so it follows that the number of steps required == LCM(number of steps required for each path)
	curs := slicesMap(cur, func(v string) int {
		s := 0
		cur := v
		for true {
			dir := path[i]
			cur = network[cur][dir]
			i = (i + 1) % len(path)
			s++
			if cur[2] == 'Z' {
				break
			}
		}
		return s
	})
	sm := 1
	for _, v := range curs {
		sm = LCM(sm, v)
	}
	fmt.Println(sm)
}
