package main

import (
	"fmt"
	"iter"
)

func main() {
	numberNames := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for v1, v2 := range Zip(SliceIter(numberNames), SliceIter(numbers)) {
		fmt.Println(v1, v2)
	}
}

func SliceIter[E any](s []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

func Zip[E, F comparable](s1 iter.Seq[E], s2 iter.Seq[F]) iter.Seq2[E, F] {
	return func(yield func(E, F) bool) {
		next1, stop1 := iter.Pull(s1)
		defer stop1()
		next2, stop2 := iter.Pull(s2)
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}
