//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"math"
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

// get roots of x^2 + ax + b = 0
func quadRoots(a, b, c float64) (float64, float64) {
	// determinant assume non-negative
	d := (b * b) - (4 * a * c)
	t1, t2 := -b/(2*a), math.Sqrt(d)/(2*a)
	return t1 - t2, t1 + t2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, times, _ := strings.Cut(scanner.Text(), ": ")
	time := slicesMap(strings.Fields(times), atoi)
	scanner.Scan()
	_, dists, _ := strings.Cut(scanner.Text(), ": ")
	dist := slicesMap(strings.Fields(dists), atoi)
	p := 1
	for i := range time {
		t, d := time[i], dist[i]
		// number of discrete positive values of i in [0...t] for which (i * (t - i)) > d holds
		// equation restated in standard form: i^2 + (-t)i + (d) = 0
		il, ih := quadRoots(1, float64(-t), float64(d))
		// since middle term of parabola is negative, it opens down
		// hence solutions for >d lie between the roots, not outside
		// calculate number of discrete values in range ([0...t] ^ [il...ih])
		r := min(t, int(math.Floor(ih))) - max(0, int(math.Ceil(il))) + 1
		if r > 0 {
			p *= r
		}
	}
	fmt.Println(p)
}
