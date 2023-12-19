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
		wflow = append(wflow, wflowarm{-1, 0, '=', ops[len(ops)-1]})
		wflows[n] = wflow
	}
	sum := 0
	for scanner.Scan() {
		str := scanner.Text()
		part := slicesMap(strings.Split(strings.Trim(str, "{}"), ","), func(s string) int { return atoi(strings.Split(s, "=")[1]) })
		cwf, cwfi := "in", 0
	machloop:
		for true {
			switch cwf {
			case "A":
				sum += part[0] + part[1] + part[2] + part[3]
				break machloop
			case "R":
				break machloop
			}
			wf := wflows[cwf][cwfi]
			f := false
			switch wf.o {
			case '=':
				f = true
			case '<':
				f = part[wf.s] < wf.n
			case '>':
				f = part[wf.s] > wf.n
			}
			if f {
				cwf, cwfi = wf.t, 0
			} else {
				cwfi++
			}
		}
	}
	fmt.Println(sum)
}
