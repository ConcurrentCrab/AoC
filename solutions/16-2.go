//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

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

var mirrors = map[byte]([][2][2]int){
	'|':  {{dirw, dirn}, {dirw, dirs}, {dire, dirn}, {dire, dirs}},
	'-':  {{dirn, dirw}, {dirn, dire}, {dirs, dirw}, {dirs, dire}},
	'/':  {{dirn, dire}, {dirw, dirs}, {dirs, dirw}, {dire, dirn}},
	'\\': {{dirn, dirw}, {dirw, dirn}, {dirs, dire}, {dire, dirs}},
}

func countAndReset(grd [][]bool) int {
	cnt := 0
	for i := range grd {
		for j := range grd[i] {
			if grd[i][j] {
				cnt++
			}
			grd[i][j] = false
		}
	}
	return cnt
}

func followBeam(srf [][]byte, energies [][]bool, pos, dir [2]int, path map[[2][2]int]bool) {
	r, c := len(srf), len(srf[0])
	for true {
		if pos[0] < 0 || pos[0] >= r || pos[1] < 0 || pos[1] >= c {
			// grid overrun
			break
		}
		if path[[2][2]int{pos, dir}] {
			// already visited
			break
		}
		path[[2][2]int{pos, dir}] = true
		energies[pos[0]][pos[1]] = true
		tile := srf[pos[0]][pos[1]]
		if tile != '.' {
			hit := false
			for _, v := range mirrors[tile] {
				if dir != v[0] {
					continue
				}
				hit = true
				followBeam(srf, energies, [2]int{pos[0] + v[1][0], pos[1] + v[1][1]}, v[1], path)
			}
			if hit {
				break
			}
		}
		pos[0], pos[1] = pos[0]+dir[0], pos[1]+dir[1]
	}
}

func main() {
	var srf [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		srf = append(srf, slices.Clone(scanner.Bytes()))
	}
	energies := gridCloneT(srf, bool(false))
	lits := []int{}
	r, c := len(srf), len(srf[0])
	for i := 0; i < r; i++ {
		followBeam(srf, energies, [2]int{i, 0}, [2]int{0, 1}, map[[2][2]int]bool{})
		lits = append(lits, countAndReset(energies))
		followBeam(srf, energies, [2]int{i, c - 1}, [2]int{0, -1}, map[[2][2]int]bool{})
		lits = append(lits, countAndReset(energies))
	}
	for j := 0; j < c; j++ {
		followBeam(srf, energies, [2]int{0, j}, [2]int{1, 0}, map[[2][2]int]bool{})
		lits = append(lits, countAndReset(energies))
		followBeam(srf, energies, [2]int{r - 1, j}, [2]int{-1, 0}, map[[2][2]int]bool{})
		lits = append(lits, countAndReset(energies))
	}
	fmt.Println(slices.Max(lits))
}
