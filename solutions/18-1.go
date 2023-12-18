//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func mathAbs[T SignedInteger](n T) T {
	if n < 0 {
		n *= -1
	}
	return n
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

var (
	dirn = [2]int{-1, 0}
	dirs = [2]int{1, 0}
	dire = [2]int{0, 1}
	dirw = [2]int{0, -1}
)

var dirsList = [4][2]int{dire, dirs, dirw, dirn}

var dirsMap = map[byte]int{
	'R': 0,
	'D': 1,
	'L': 2,
	'U': 3,
}

func main() {
	var lagoon = [][2]int{{0, 0}}
	var perim = 0
	var cur = lagoon[0]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		strf := strings.Fields(str)
		d, n := dirsMap[strf[0][0]], atoi(strf[1])
		perim += n
		dir := dirsList[d]
		cur[0], cur[1] = cur[0]+(dir[0]*n), cur[1]+(dir[1]*n)
		lagoon = append(lagoon, cur)
	}
	area := 0
	// we can ignore the last vertex pair as it should be the same as the first
	for i := 0; i < len(lagoon)-1; i++ {
		lagc, lagn := lagoon[i], lagoon[i+1]
		area += (lagc[0] + lagn[0]) * (lagc[1] - lagn[1])
	}
	farea := ((mathAbs(area) + perim) / 2) + 1
	fmt.Println(farea)
}
