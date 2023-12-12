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

func totalPoints(w, h []int) int {
	s := 0
	slices.Sort(w)
	slices.Sort(h)
	iw, ih := 0, 0
	for iw < len(w) && ih < len(h) {
		if w[iw] == h[ih] {
			s++
			iw++
			ih++
		} else if w[iw] < h[ih] {
			iw++
		} else if h[ih] < w[iw] {
			ih++
		}
	}
	return s
}

func main() {
	var cards []int
	var won []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		_, a, _ := strings.Cut(str, ": ")
		ws, hs, _ := strings.Cut(a, " | ")
		w := slicesMap(strings.Fields(ws), atoi)
		h := slicesMap(strings.Fields(hs), atoi)
		p := totalPoints(w, h)
		cards = append(cards, p)
	}
	for i := range cards {
		won = append(won, i)
	}
	// won grows during iteration so this runs as long as our luck does
	for i := 0; i < len(won); i++ {
		p := cards[won[i]]
		for j := 1; j <= p; j++ {
			won = append(won, won[i]+j)
		}
	}
	fmt.Println(len(won))
}
