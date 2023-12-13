//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func gridFindSymmetry[T comparable](grd [][]T) (int, int) {
	r, c := len(grd), len(grd[0])
	for i := 0; i < (r - 1); i++ {
		imp := 0
		for o := 0; o < min(i+1, r-i-1); o++ {
			ios, ioe := i-o, i+o+1
			for j := 0; j < c; j++ {
				if grd[ios][j] != grd[ioe][j] {
					imp++
				}
			}
		}
		if imp == 1 {
			return i, -1
		}
	}
	for j := 0; j < (c - 1); j++ {
		imp := 0
		for o := 0; o < min(j+1, c-j-1); o++ {
			jos, joe := j-o, j+o+1
			for i := 0; i < r; i++ {
				if grd[i][jos] != grd[i][joe] {
					imp++
				}
			}
		}
		if imp == 1 {
			return -1, j
		}
	}
	return -1, -1
}

func main() {
	var mirr [][]byte
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		// flush mappings on a blank line
		// NOTE: because of this heuristic, input has to have 2 newlines at end
		if str == "" {
			h, v := gridFindSymmetry(mirr)
			if h > -1 {
				sum += 100 * (h + 1)
			}
			if v > -1 {
				sum += 1 * (v + 1)
			}
			mirr = mirr[:0]
			continue
		}
		mirr = append(mirr, []byte(str))
	}
	fmt.Println(sum)
}
