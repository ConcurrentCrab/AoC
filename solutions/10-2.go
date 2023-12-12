//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

// find values f accepts in kernel around i, j
func gridFindInRadiusCoords[T any](grd [][]T, krnl [][2]int, i, j int, f func(T, int, int) bool) [][2]int {
	var res [][2]int
	r, c := len(grd), len(grd[0])
	for _, o := range krnl {
		io, jo := i+o[0], j+o[1]
		if io < 0 || io >= r || jo < 0 || jo >= c {
			continue
		}
		if f(grd[io][jo], io, jo) {
			res = append(res, [2]int{io, jo})
		}
	}
	return res
}

func gridCrawlFromEdgesRecr[T any](grd [][]T, krnl [][2]int, i, j int, f func([][]T, int, int) bool) {
	if f(grd, i, j) {
		return
	}
	r, c := len(grd), len(grd[0])
	for _, o := range krnl {
		io, jo := i+o[0], j+o[1]
		if io < 0 || io >= r || jo < 0 || jo >= c {
			continue
		}
		gridCrawlFromEdgesRecr(grd, krnl, io, jo, f)
	}
}

func gridCrawlFromEdges[T any](grd [][]T, krnl [][2]int, f func([][]T, int, int) bool) {
	r, c := len(grd), len(grd[0])
	for i := 0; i < r; i++ {
		gridCrawlFromEdgesRecr(grd, krnl, i, 0, f)
		gridCrawlFromEdgesRecr(grd, krnl, i, c-1, f)
	}
	for j := 0; j < c; j++ {
		gridCrawlFromEdgesRecr(grd, krnl, 0, j, f)
		gridCrawlFromEdgesRecr(grd, krnl, r-1, j, f)
	}
}

// Make another grid of same dimensions but with given type
func gridCloneT[T, U any](a [][]T, _ U) [][]U {
	r, c := len(a), len(a[0])
	ac := make([][]U, r)
	for i := range ac {
		ac[i] = make([]U, c)
	}
	return ac
}

var (
	dirn = [2]int{-1, 0}
	dirs = [2]int{1, 0}
	dire = [2]int{0, 1}
	dirw = [2]int{0, -1}
)

var dirsList = [4][2]int{dirn, dirs, dire, dirw}

var pipeEnds = map[byte]([2][2]int){
	'|': {dirn, dirs},
	'-': {dire, dirw},
	'L': {dirn, dire},
	'J': {dirn, dirw},
	'7': {dirw, dirs},
	'F': {dire, dirs},
}

var pipeGrown = map[byte]([3][3]byte){
	'S': {
		{' ', '|', ' '},
		{'-', 'S', '-'},
		{' ', '|', ' '},
	},
	'|': {
		{' ', '|', ' '},
		{' ', '|', ' '},
		{' ', '|', ' '},
	},
	'-': {
		{' ', ' ', ' '},
		{'-', '-', '-'},
		{' ', ' ', ' '},
	},
	'L': {
		{' ', '|', ' '},
		{' ', 'L', '-'},
		{' ', ' ', ' '},
	},
	'J': {
		{' ', '|', ' '},
		{'-', 'J', ' '},
		{' ', ' ', ' '},
	},
	'7': {
		{' ', ' ', ' '},
		{'-', '7', ' '},
		{' ', '|', ' '},
	},
	'F': {
		{' ', ' ', ' '},
		{' ', 'F', '-'},
		{' ', '|', ' '},
	},
}

func main() {
	var srf, srfa [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		srf = append(srf, slices.Clone(scanner.Bytes()))
	}
	// blow up grid to 3x and stitch back the blown up tiles
	for i := range srf {
		var lines [3][]byte
		for j := range srf[i] {
			cg := pipeGrown[srf[i][j]]
			for l := range lines {
				lines[l] = append(lines[l], cg[l][:]...)
			}
		}
		for l := range lines {
			srfa = append(srfa, lines[l])
		}
	}
	srf = srfa
	dists := gridCloneT(srf, int(0))
	pos := gridFind(srf, func(b byte) bool { return b == 'S' })[0]
	dists[pos[0]][pos[1]] = 1
	paths := gridFindInRadiusCoords(srf, dirsList[:], pos[0], pos[1], func(b byte, i, j int) bool {
		f := false
		for _, conv := range pipeEnds[b] {
			if pos == [2]int{i + conv[0], j + conv[1]} {
				f = true
				break
			}
		}
		return f
	})
	for _, v := range paths {
		prev, cur := pos, v
		dir := [2]int{prev[0] - cur[0], prev[1] - cur[1]}
		steps := 1
		for cur != pos {
			dists[cur[0]][cur[1]] = 1
			piece := srf[cur[0]][cur[1]]
			m := -1
			for n, end := range pipeEnds[piece] {
				if end == dir {
					m = n
				}
			}
			if m >= 0 {
				// choose the other end
				oend := pipeEnds[piece][1-m]
				prev[0], prev[1] = cur[0], cur[1]
				cur[0], cur[1] = cur[0]+oend[0], cur[1]+oend[1]
				dir[0], dir[1] = prev[0]-cur[0], prev[1]-cur[1]
				steps++
			} else {
				// dead end
				break
			}
		}
	}
	// flood fill
	gridCrawlFromEdges(dists, dirsList[:], func(grd [][]int, i, j int) bool {
		if grd[i][j] == 1 || grd[i][j] == 2 {
			return true
		}
		grd[i][j] = 2
		return false
	})
	// shrink back grid 3x
	distss := make([][]int, len(dists)/3)
	for i := 1; i < len(dists); i += 3 {
		rs := make([]int, len(dists[0])/3)
		for j := 1; j < len(dists[0]); j += 3 {
			rs[j/3] = dists[i][j]
		}
		distss[i/3] = rs
	}
	dists = distss
	dmax := gridFind(dists, func(n int) bool { return n == 0 })
	fmt.Println(len(dmax))
}
