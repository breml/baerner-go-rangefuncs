package main

import (
	"fmt"
	"iter"
)

func main() {
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for _, v := range Backwards(numbers) {
		fmt.Println(v)
	}
}

func Backwards[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}
