package main

import (
	"context"
	"fmt"
)

func Load(ctx context.Context) {
	fmt.Println("plugin2 is loaded.")
	for range 10 {
		fmt.Println("hello world from plugin 2")
	}

	for {
		select {
		case <-ctx.Done():
			// Context canceled, clean up and exit
			fmt.Println("plugin 2 received cancellation signal. Exiting...\n")
			return
		}
	}
}
