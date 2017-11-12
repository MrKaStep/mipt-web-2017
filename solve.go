package main

import (
	"strings"
	"unicode"
)

// RemoveEven removes even numbers form slice
func RemoveEven(s []int) []int {
	result := make([]int, 0)
	for _, a := range s {
		if a%2 == 1 {
			result = append(result, a)
		}
	}
	return result
}

// PowerGenerator returns generator of powers of given integer
func PowerGenerator(a int) func() int {
	t := 1
	gen := func() int {
		t *= 2
		return t
	}
	return gen
}

// DifferentWordsCount return number of different words in given string
func DifferentWordsCount(s string) int {
	w := make([]string, 0)
	m := make(map[string]int)
	for _, c := range s {
		if unicode.IsLetter(rune(c)) {
			w = append(w, string(unicode.ToLower(rune(c))))
		} else {
			if len(w) > 0 {
				m[strings.Join(w, "")]++
				w = make([]string, 0)
			}
		}
	}
	return len(m)
}

// func main() {
// 	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
// }
