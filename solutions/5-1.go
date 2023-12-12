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

func main() {
	var seeds, seedsn, seedsl []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, s, _ := strings.Cut(scanner.Text(), ": ")
	seeds = slicesMap(strings.Fields(s), atoi)
	for scanner.Scan() {
		str := scanner.Text()
		// flush mappings on a blank line
		// NOTE: because of this heuristic, input has to have 2 newlines at end
		if str == "" {
			seeds = append(seeds, seedsn...)
			seedsn = seedsn[:0]
			scanner.Scan() // skip header line
			continue
		}
		nmap := [3]int(slicesMap(strings.Fields(str), atoi))
		for _, v := range seeds {
			if v >= nmap[1] && v < (nmap[1]+nmap[2]) {
				seedsn = append(seedsn, v+nmap[0]-nmap[1])
			} else {
				seedsl = append(seedsl, v)
			}
		}
		seeds, seedsl = seedsl, seeds[:0]
	}
	fmt.Println(slices.Min(seeds))
}
