package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func startProducers(ctx context.Context, wg *sync.WaitGroup, productionLines []productionLine) []chan string {
	productionLineDeliveries := make([]chan string, 0, len(productionLines))
	wg.Add(len(productionLines))

	for i := 0; i < len(productionLines); i++ {
		i := i
		delivery := make(chan string, 0)
		productionLineDeliveries = append(productionLineDeliveries, delivery)
		go producer(ctx, wg, delivery, i, productionLines[i].kind, productionLines[i].items, productionLines[i].productionTime)
	}

	return productionLineDeliveries
}

func producer(ctx context.Context, wg *sync.WaitGroup, delivery chan string, idx int, kind string, items int, productionTime time.Duration) {
	fmt.Printf("📣 start producer %d for %q\n", idx, kind)
	defer wg.Done()
	defer close(delivery)
	defer fmt.Printf("📣 producer %d for %q done\n", idx, kind)

	t := time.NewTicker(productionTime)
	defer t.Stop()

	for i := 0; i < items; i++ {
		select {
		case <-t.C:
			// item produced
		case <-ctx.Done():
			// today, we go home earlier
			return
		}

		select {
		case delivery <- fmt.Sprintf("%s (%d)", kind, i):
			// deliver the thing
		case <-ctx.Done():
			// today, we go home earlier
			return
		}

	}
}
