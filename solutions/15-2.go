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

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func hashStr(s string) int {
	b := []byte(s) // iterating over a string returns runes not bytes
	n := 0
	for _, v := range b {
		n += int(v)
		n *= 17
		n &= (1 << 8) - 1 // mod 256
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	type lens struct {
		lbl string
		pow int
	}
	var boxes [256][]lens
	for _, s := range strings.Split(str, ",") {
		l, o, p := "", byte(0), 0
		switch {
		case s[len(s)-1] == '-':
			l, o = s[:len(s)-1], s[len(s)-1]
		case s[len(s)-2] == '=':
			l, o, p = s[:len(s)-2], s[len(s)-2], atoi(s[len(s)-1:])
		}
		h := hashStr(l)
		i := slices.IndexFunc(boxes[h], func(v lens) bool { return v.lbl == l })
		switch {
		case o == '=' && i > -1:
			boxes[h][i].pow = p
		case o == '=':
			boxes[h] = append(boxes[h], lens{l, p})
		case o == '-' && i > -1:
			boxes[h] = slices.Delete(boxes[h], i, i+1)
		}
	}
	pow := 0
	for i, b := range boxes {
		for j, l := range b {
			pow += (i + 1) * (j + 1) * (l.pow)
		}
	}
	fmt.Println(pow)
}
