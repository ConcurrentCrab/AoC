//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func findCycle[T comparable](a []T) int {
	c := 1
	for c < len(a) {
		f := true
		for ci := 0; ci < c; ci++ {
			for cj := ci + c; cj < len(a); cj += c {
				if a[ci] != a[cj] {
					f = false
					goto endloop
				}
			}
		}
	endloop:
		if f {
			return c
		}
		c++
	}
	return 0
}

func rotationCycle(grd [][]byte) {
	r, c := len(grd), len(grd[0])
	// north tilt
	for j := 0; j < c; j++ {
		hold := -1
		for i := 0; i < r; i++ {
			switch grd[i][j] {
			case '#':
				hold = i
			case 'O':
				hold++
				grd[i][j] = '.'
				grd[hold][j] = 'O'
			}
		}
	}
	// west tilt
	for i := 0; i < r; i++ {
		hold := -1
		for j := 0; j < c; j++ {
			switch grd[i][j] {
			case '#':
				hold = j
			case 'O':
				hold++
				grd[i][j] = '.'
				grd[i][hold] = 'O'
			}
		}
	}
	// south tilt
	for j := 0; j < c; j++ {
		hold := r
		for i := r - 1; i >= 0; i-- {
			switch grd[i][j] {
			case '#':
				hold = i
			case 'O':
				hold--
				grd[i][j] = '.'
				grd[hold][j] = 'O'
			}
		}
	}
	// east tilt
	for i := 0; i < r; i++ {
		hold := c
		for j := c - 1; j >= 0; j-- {
			switch grd[i][j] {
			case '#':
				hold = j
			case 'O':
				hold--
				grd[i][j] = '.'
				grd[i][hold] = 'O'
			}
		}
	}
}

func findLoad(grd [][]byte) int {
	load := 0
	r, c := len(grd), len(grd[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grd[i][j] == 'O' {
				load += r - i
			}
		}
	}
	return load
}

func main() {
	var srf [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		srf = append(srf, slices.Clone(scanner.Bytes()))
	}
	ns := []int{}
	for n := 0; n < 2000; n++ { // ?? WTF? MAGIC NUMBER VOODOO?
		ns = append(ns, findLoad(srf))
		rotationCycle(srf)
	}
	mgc := findCycle(ns[1000:]) // ???? WTFFF! THIS IS FRANKLY SHAMEFUL CODE
	mgci := ((1000 / mgc) * mgc) + (1000000000 % mgc)
	fmt.Println(ns[mgci])
}
