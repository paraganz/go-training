package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var out io.Writer = os.Stdout

func letters(s string) map[rune]int {
	r := []rune(s)
	var myMap = make(map[rune]int)
	for i := 0; i < len(r); i++ {
		for j := 1; j < len(r); j++ {
			if r[i] == r[j] {
				myMap[r[i]] = 0
			}
		}
	}

	for k := range myMap {
		for i := 0; i < len(r); i++ {
			if r[i] == k {
				myMap[r[i]] = myMap[r[i]] + 1
			}
		}
	}
	return myMap
}

func sortLetters(m map[rune]int) []string {
	keys := []int{}
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	sortedKeys := make([]string, len(keys))
	for pos, k := range keys {
		sortedKeys[pos] = fmt.Sprintf("%c:%d", k, m[rune(k)])
	}
	return sortedKeys
}

func main() {
	fmt.Println(strings.Join(sortLetters(letters("aba")), "\n"))
}
