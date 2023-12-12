//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func slicesMap[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func mathAbs[T SignedInteger](n T) T {
	if n < 0 {
		n *= -1
	}
	return n
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

const exprate = 1000000 - 1

func main() {
	var space [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		space = append(space, slices.Clone(scanner.Bytes()))
	}
	mtc := gridFind(space, func(b byte) bool { return b == '#' })
	// get list of rows and columns that contain stars, and sort and uniq them
	mtcrc := [2][]int{
		slicesMap(mtc, func(v [2]int) int { return v[0] }),
		slicesMap(mtc, func(v [2]int) int { return v[1] }),
	}
	for i := range mtcrc {
		slices.Sort(mtcrc[i])
		mtcrc[i] = slices.Compact(mtcrc[i])
	}
	for i := range mtc {
		// shift down/right by rate * no. of empty rows/columns before it
		// no of empty r/c before = original r/c - index of r/c in matches list
		mtc[i][0] += exprate * (mtc[i][0] - slices.Index(mtcrc[0], mtc[i][0]))
		mtc[i][1] += exprate * (mtc[i][1] - slices.Index(mtcrc[1], mtc[i][1]))
	}
	sum := 0
	for i := 0; i < len(mtc); i++ {
		for j := i + 1; j < len(mtc); j++ {
			dist := mathAbs(mtc[i][0]-mtc[j][0]) + mathAbs(mtc[i][1]-mtc[j][1])
			sum += dist
		}
	}
	fmt.Println(sum)
}
