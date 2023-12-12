//go:build ignore

package main

import (
	"bufio"
	"cmp"
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

func slicesUniqCount[T comparable](a []T) map[T]int {
	r := make(map[T]int)
	for _, v := range a {
		_, e := r[v]
		if !e {
			r[v] = 0
		}
		r[v] += 1
	}
	return r
}

// homegrow maps.Values() impl until it's stabilised
func mapsValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, len(m))
	i := 0
	for _, v := range m {
		r[i] = v
		i++
	}
	return r
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var cardRanks = []byte("J23456789TQKA")

func cardRank(card byte) int {
	for i := range cardRanks {
		if card == cardRanks[i] {
			return i
		}
	}
	return -1
}

func handRank(cards [5]byte) int {
	uniqm := slicesUniqCount(cards[:])
	// count and remove jokers
	jn := uniqm['J']
	uniqm['J'] = 0
	uniqs := mapsValues(uniqm)
	slices.Sort(uniqs)
	slices.Reverse(uniqs)
	// add jokers back to biggest group
	uniqs[0] += jn
	switch {
	case uniqs[0] == 1:
		// high card
		return 0
	case uniqs[0] == 2 && uniqs[1] == 1:
		// one pair
		return 1
	case uniqs[0] == 2 && uniqs[1] == 2:
		// two pair
		return 2
	case uniqs[0] == 3 && uniqs[1] == 1:
		// three of a kind
		return 3
	case uniqs[0] == 3 && uniqs[1] == 2:
		// full house
		return 4
	case uniqs[0] == 4:
		// four of a kind
		return 5
	case uniqs[0] == 5:
		// five of a kind
		return 6
	}
	return -1
}

func main() {
	var hands [][7]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		h, b, _ := strings.Cut(str, " ")
		hb := [5]byte([]byte(h))
		hand := [7]int{atoi(b), handRank(hb)}      // bid, rank of hand
		copy(hand[2:], slicesMap(hb[:], cardRank)) // rank of cards
		hands = append(hands, hand)
	}
	slices.SortStableFunc(hands, func(a, b [7]int) int {
		r := cmp.Compare(a[1], b[1]) // compare hand rank
		if r == 0 {
			r = slices.Compare(a[2:], b[2:]) // compare card rank
		}
		return r
	})
	s := 0
	for i, v := range hands {
		s += v[0] * (i + 1)
	}
	fmt.Println(s)
}
