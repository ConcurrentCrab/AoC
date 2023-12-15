//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	var srf [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		srf = append(srf, slices.Clone(scanner.Bytes()))
	}
	r, c := len(srf), len(srf[0])
	for j := 0; j < c; j++ {
		hold := -1
		for i := 0; i < r; i++ {
			switch srf[i][j] {
			case '#':
				hold = i
			case 'O':
				hold++
				srf[i][j] = '.'
				srf[hold][j] = 'O'
			}
		}
	}
	load := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if srf[i][j] == 'O' {
				load += r - i
			}
		}
	}
	fmt.Println(load)
}
