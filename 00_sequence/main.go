package main

import (
	"fmt"
	"iter"
)

func main() {
	seq := Sequence()
	for i := range seq {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

func Sequence() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			cont := yield(i)
			if !cont {
				return
			}
		}
	}
}
