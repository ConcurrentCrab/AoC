//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Regexp.FindAllIndex only returns non-overlapping matches, so this does that job
func regexpFindAllIndexOverlapping(r *regexp.Regexp, s string) [][2]int {
	var ma [][2]int
	i := 0
	for i < len(s) {
		m := r.FindStringIndex(s[i:])
		if m == nil {
			break
		}
		ma = append(ma, [2]int{i + m[0], i + m[1]})
		i += m[0] + 1
	}
	return ma
}

func strtonum(s string) int {
	if s[0] >= '0' && s[0] <= '9' {
		return int(s[0] - '0')
	}
	v := -1
	switch s {
	case "zero":
		v = 0
	case "one":
		v = 1
	case "two":
		v = 2
	case "three":
		v = 3
	case "four":
		v = 4
	case "five":
		v = 5
	case "six":
		v = 6
	case "seven":
		v = 7
	case "eight":
		v = 8
	case "nine":
		v = 9
	}
	return v
}

var rgx = regexp.MustCompile("([0-9]|zero|one|two|three|four|five|six|seven|eight|nine)")

func main() {
	s := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		m := regexpFindAllIndexOverlapping(rgx, str)
		a, b := m[0], m[len(m)-1]
		s += strtonum(str[a[0]:a[1]]) * 10
		s += strtonum(str[b[0]:b[1]])
	}
	fmt.Println(s)
}
