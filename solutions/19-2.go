//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
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

var partMap = map[byte]int{
	'x': 0,
	'm': 1,
	'a': 2,
	's': 3,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	type wflowarm struct {
		s, n int
		o    byte
		t    string
	}
	var wflows = make(map[string]([]wflowarm))
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		n, str, _ := strings.Cut(str, "{")
		ops := strings.Split(strings.Trim(str, "{}"), ",")
		wflow := slicesMap(ops[:len(ops)-1], func(str string) wflowarm {
			s, o, r := str[0], str[1], str[2:]
			n, t, _ := strings.Cut(r, ":")
			return wflowarm{partMap[s], atoi(n), o, t}
		})
		wflow = append(wflow, wflowarm{0, 0, '=', ops[len(ops)-1]})
		wflows[n] = wflow
	}
	type partrange struct {
		rng  [4][2]int
		cwf  string
		cwfi int
	}
	var parts, partsn []partrange
	parts = []partrange{{[4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, "in", 0}}
	sum := 0
	for len(parts) > 0 {
		for _, v := range parts {
			switch v.cwf {
			case "A":
				sum += v.rng[0][1] * v.rng[1][1] * v.rng[2][1] * v.rng[3][1]
				continue
			case "R":
				continue
			}
			wf := wflows[v.cwf][v.cwfi]
			var srng, frng [][2]int
			switch wf.o {
			case '=':
				srng = [][2]int{v.rng[wf.s]}
			case '<':
				frng, srng, _ = intersectRanges(v.rng[wf.s], [2]int{1, wf.n - 1})
			case '>':
				frng, srng, _ = intersectRanges(v.rng[wf.s], [2]int{wf.n + 1, 4000 - wf.n})
			}
			v.cwfi++
			for _, r := range frng {
				v.rng[wf.s] = r
				partsn = append(partsn, v)
			}
			v.cwf, v.cwfi = wf.t, 0
			for _, r := range srng {
				v.rng[wf.s] = r
				partsn = append(partsn, v)
			}
		}
		parts, partsn = partsn, parts[:0]
	}
	fmt.Println(sum)
}
