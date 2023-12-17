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
	dirw = [2]int{0, -1}
	dirs = [2]int{1, 0}
	dire = [2]int{0, 1}
)

var dirsList = [4][2]int{dirn, dirw, dirs, dire}

func walkCrucible(srf [][]byte) int {
	r, c := len(srf), len(srf[0])
	heats := gridCloneT(srf, [4][11]int{})
	heatr := []int{}
	var cstates, nstates [][5]int
	cstates = [][5]int{{0, 1, 3, 0, 1}, {1, 0, 2, 0, 1}}
	for len(cstates) > 0 {
	stateloop:
		for i := range cstates {
			pi, pj, dir, heat, cont := cstates[i][0], cstates[i][1], cstates[i][2], cstates[i][3], cstates[i][4]
			if pi < 0 || pi >= r || pj < 0 || pj >= c {
				// grid overrun
				continue stateloop
			}
			heat += int(srf[pi][pj] - '0')
			if pi == (r-1) && pj == (c-1) && cont >= 4 {
				// reached bottom left
				heatr = append(heatr, heat)
				continue stateloop
			}
			for c, h := range heats[pi][pj][dir] {
				if c >= 4 && c <= cont && h != 0 && h <= heat {
					// inefficient path
					continue stateloop
				}
			}
			heats[pi][pj][dir][cont] = heat
			for d := range dirsList {
				c := 1
				switch {
				case d == ((dir + 2) % 4):
					// no turning back
					continue
				case d == dir && cont == 10:
					// no more going straight
					continue
				case d == dir:
					c = cont + 1
				case d != dir && cont < 4:
					// no turning
					continue
				}
				nstates = append(nstates, [5]int{pi + dirsList[d][0], pj + dirsList[d][1], d, heat, c})
			}
		}
		cstates, nstates = nstates, cstates[:0]
	}
	return slices.Min(heatr)
}

func main() {
	var srf [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		srf = append(srf, slices.Clone(scanner.Bytes()))
	}
	ep := walkCrucible(srf)
	fmt.Println(ep)
}
