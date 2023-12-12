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

// take two intervals, return arrays of intervals of their intersection and leftovers
func intersectRanges(a, b [2]int) (al, il, bl [][2]int) {
	as, ae, bs, be := a[0], a[0]+a[1], b[0], b[0]+b[1]
	if (ae < bs) || (be < as) {
		al = [][2]int{a}
		bl = [][2]int{b}
		return
	}
	is, ie := max(as, bs), min(ae, be)
	il = [][2]int{{is, ie - is}}
	if as < is {
		al = append(al, [2]int{as, is - as})
	} else if bs < is {
		bl = append(bl, [2]int{bs, is - bs})
	}
	if ie < ae {
		al = append(al, [2]int{ie, ae - ie})
	} else if ie < be {
		bl = append(bl, [2]int{ie, be - ie})
	}
	return
}

func main() {
	var seeds, seedsn, seedsl [][2]int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, s, _ := strings.Cut(scanner.Text(), ": ")
	seedi := slicesMap(strings.Fields(s), atoi)
	for i := 0; i < len(seedi); i += 2 {
		seeds = append(seeds, [2]int{seedi[i], seedi[i+1]})
	}
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
			l, o, _ := intersectRanges(v, [2]int(nmap[1:]))
			for _, w := range o {
				w[0] += nmap[0] - nmap[1]
				seedsn = append(seedsn, w)
			}
			seedsl = append(seedsl, l...)
		}
		seeds, seedsl = seedsl, seeds[:0]
	}
	fmt.Println(slices.Min(slicesMap(seeds, func(a [2]int) int { return a[0] })))
}
