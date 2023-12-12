//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// keep differentiating until you find a flat line of 0 and extrapolate back
func extrapolateByDifferentiation[T Integer](a []T) (T, T) {
	a = slices.Clone(a)
	var edge [][2]T
	for slices.ContainsFunc(a, func(v T) bool { return v != 0 }) {
		edge = append(edge, [2]T{a[0], a[len(a)-1]})
		for i := range a[:len(a)-1] {
			a[i] = a[i+1] - a[i]
		}
		a = a[:len(a)-1]
	}
	p, n := T(0), T(0)
	for len(edge) > 0 {
		e := edge[len(edge)-1]
		p, n = e[0]-p, e[1]+n
		edge = edge[:len(edge)-1]
	}
	return p, n
}

func main() {
	s := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		vals := slicesMap(strings.Fields(str), atoi)
		p, _ := extrapolateByDifferentiation(vals)
		s += p
	}
	fmt.Println(s)
}
