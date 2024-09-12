package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type productionLine struct {
	kind           string
	items          int
	productionTime time.Duration
}

var productionLines = []productionLine{
	{
		kind:           "train",
		items:          1,
		productionTime: 3000 * time.Millisecond,
	},
	{
		kind:           "car",
		items:          3,
		productionTime: 1000 * time.Millisecond,
	},
	{
		kind:           "bike",
		items:          10,
		productionTime: 350 * time.Millisecond,
	},
	// {
	// 	kind:           "gopher sticker",
	// 	items:          20,
	// 	productionTime: 150 * time.Millisecond,
	// },
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	productionLineDeliveries := startProducers(ctx, wg, productionLines)

	for item := range ConsumeWithContext(ctx, productionLineDeliveries...) {
		fmt.Println("✅", item)
		// if item == "gopher sticker (12)" {
		// 	fmt.Println("produced gophers (12), finish consuming early")
		// 	break
		// }

		// if item == "bike (6)" {
		// 	fmt.Println("produced bike (6), finish consuming early")
		// 	cancel()
		// }
	}

	fmt.Println("📣 consume loop done")

	cancel()

	wg.Wait()

	fmt.Println("📣 all done")
}
