package main

import (
	"context"
	"fmt"
)

func Load(ctx context.Context) {
	fmt.Println("plugin1 is loaded.")
	for range 10 {
		fmt.Println("hello world from plugin 1")
	}

	<-ctx.Done()
}
