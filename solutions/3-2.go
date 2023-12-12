//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

// run f for all elements and return matches
func gridFind[T any](grd [][]T, f func(T) bool) [][2]int {
	var res [][2]int
	for i, r := range grd {
		for j, c := range r {
			if f(c) {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
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

func main() {
	s := 0
	var sch [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sch = append(sch, slices.Clone(scanner.Bytes()))
	}
	stars := gridFind(sch, func(b byte) bool { return b == '*' })
	for _, star := range stars {
		numr := gridFindInRadius(sch, star[0], star[1], isDigit)
		numm := slicesMap(numr, func(n [2]int) [3]int {
			mtc := gridFindHRun(sch, n[0], n[1], isDigit)
			return [3]int{n[0], mtc[0], mtc[1]}
		})
		numm = slices.Compact(numm)
		if len(numm) != 2 {
			continue
		}
		nums := slicesMap(numm, func(n [3]int) int {
			return atoi(string(sch[n[0]][n[1]:n[2]]))
		})
		s += nums[0] * nums[1]
	}
	fmt.Println(s)
}
