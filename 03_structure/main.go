package main

import (
	"context"
	"iter"
	"time"

	"github.com/fatih/color"
)

func Iterator[E any](ctx context.Context, s []E) iter.Seq2[int, E] {
	c := color.New(color.BgHiYellow, color.FgHiBlack)
	c.EnableColor()
	c.Println("  🛫 ① global initialization of iterator")
	i := 0

	context.AfterFunc(ctx, func() {
		c.Println("  🛬 ⑧ global cleanup of iterator")
	})

	return func(yield func(int, E) bool) {
		c := color.New()
		c.EnableColor()

		// Setup, initialize iterator
		c.Println("    🛫 ② setup iterator")

		defer func() {
			// Teardown iterator, cleanup
			c.Println("    🛬 ⑦ teardown iterator")
		}()

		for i < len(s) {
			// before each iteration
			c.Println("      ➡ ③ before iteration, idx:", i)

			cont := yield(i, s[i])

			// after each iteration
			c.Println("      ⬅ ⑤ after iteration, idx:", i)
			i++

			if !cont {
				// before breaking
				c.Println("      ⏎ ⑥ before cancel of iterator (external), idx:", i)
				return
			}
		}

		// before regular end of iterator
		c.Println("      ⏎ ⑥ before end of iterator (internal), idx:", i)
	}
}

func main() {
	c := color.New(color.BgHiGreen, color.FgHiBlack)
	c.EnableColor()

	c.Println("▶️  main start")

	defer func() {
		c.Println("⏹️  main end")
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(10 * time.Millisecond)
	}()

	iterator := Iterator(ctx, []int{2, 3, 5, 7, 11, 13})

	for i, v := range iterator {
		if i == 2 {
			c.Println("        ⏭️  continue")
			continue
		}

		c.Println(c.Sprintf("        ➰ ④ in loop body, idx: %d, value: %d", i, v))

		if v == 11 {
			c.Println("        🛑 request break")
			break
		}
	}

	c.Println("⏩  main between iterators")

	for i, v := range iterator {
		c.Println(c.Sprintf("        ➰ ④ in loop body, idx: %d, value: %d", i, v))
	}
}
