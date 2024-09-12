package main

import (
	"cmp"
	"fmt"
	"iter"
	"sort"
)

func main() {
	myMap := map[string]string{
		"4": "four",
		"2": "two",
		"3": "three",
		"6": "six",
		"1": "one",
		"5": "five",
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	fmt.Println("=== Ordered by key ===")

	for k, v := range OrderedByKey(myMap) {
		fmt.Println(k, v)
	}
}

func OrderedByKey[K cmp.Ordered, E any](m map[K]E) iter.Seq2[K, E] {
	return func(yield func(K, E) bool) {
		keys := make([]K, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		for _, k := range keys {
			if !yield(k, m[k]) {
				return
			}
		}
	}
}
