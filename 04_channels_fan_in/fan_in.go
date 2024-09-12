package main

import (
	"context"
	"fmt"
	"iter"
	"sync"
)

func Consume[E any](chans ...chan E) iter.Seq[E] {
	return ConsumeWithContext(context.Background(), chans...)
}

func ConsumeWithContext[E any](ctx context.Context, chans ...chan E) iter.Seq[E] {
	return func(yield func(E) bool) {
		res := make(chan E, 0)
		defer close(res)
		done := make(chan bool, len(chans))
		defer close(done)

		wg := sync.WaitGroup{}
		wg.Add(len(chans))
		defer wg.Wait()

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for i, c := range chans {
			go func(c chan E) {
				fmt.Printf("📣 start consume chan %d\n", i)

				defer wg.Done()

				defer fmt.Printf("📣 consume for chan %d done\n", i)

				for {
					select {
					case v, ok := <-c:
						if !ok {
							done <- true
							return
						}
						select {
						case res <- v:
						case <-ctx.Done():
							return
						}

					case <-ctx.Done():
						return
					}
				}
			}(c)
		}

		activeChans := len(chans)

		defer fmt.Println("📣 consume iterator done")

		for {
			select {
			case v := <-res:
				if !yield(v) {
					return
				}
			case <-done:
				activeChans--
				if activeChans == 0 {
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}
}
