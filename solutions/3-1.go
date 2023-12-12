//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

// convolution kernel
var gridFindInRadiusKernel = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

// find values f accepts in kernel around i, j
func gridFindInRadius[T any](grd [][]T, i, j int, f func(T) bool) [][2]int {
	var res [][2]int
	r, c := len(grd), len(grd[0])
	for _, o := range gridFindInRadiusKernel {
		io, jo := i+o[0], j+o[1]
		if io < 0 || io >= r || jo < 0 || jo >= c {
			continue
		}
		if f(grd[io][jo]) {
			res = append(res, [2]int{io, jo})
		}
	}
	return res
}

// find horizontal runs of values that f accepts around i, j
func gridFindHRun[T any](grd [][]T, i, j int, f func(T) bool) [2]int {
	js, je := j, j+1
	for ji := js - 1; ji >= 0 && f(grd[i][ji]); ji-- {
		js--
	}
	for ji := je; ji < len(grd[i]) && f(grd[i][ji]); ji++ {
		je++
	}
	return [2]int{js, je}
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	switch {
	case b >= '0' && b <= '9':
	case b == '.':
	default:
		return true
	}
	return false
}

func main() {
	s := 0
	var sch [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sch = append(sch, slices.Clone(scanner.Bytes()))
	}
	nums := [][3]int{}
	for i, r := range sch {
		j := 0
		for j < len(r) {
			if !isDigit(r[j]) {
				j++
				continue
			}
			mtc := gridFindHRun(sch, i, j, isDigit)
			nums = append(nums, [3]int{i, mtc[0], mtc[1]})
			j = mtc[1]
		}
	}
	nums = slicesFilter(nums, func(n [3]int) bool {
		for j := n[1]; j < n[2]; j++ {
			sym := gridFindInRadius(sch, n[0], j, isSymbol)
			if len(sym) > 0 {
				return true
			}
		}
		return false
	})
	for _, num := range nums {
		s += atoi(string(sch[num[0]][num[1]:num[2]]))
	}
	fmt.Println(s)
}
