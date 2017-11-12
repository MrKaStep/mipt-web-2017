package main

import (
	"fmt"
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
	d := a
	gen := func() int {
		t *= d
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

func main() {
	gen2 := PowerGenerator(2)
	gen3 := PowerGenerator(3)
	fmt.Println(gen2())
	fmt.Println(gen2())
	fmt.Println(gen2())
	fmt.Println(gen3())
	fmt.Println(gen3())
	fmt.Println(gen3())
	fmt.Println(gen2())
	fmt.Println(gen2())
	fmt.Println(gen2())
	fmt.Println(gen3())
	fmt.Println(gen3())
	fmt.Println(gen3())
}
