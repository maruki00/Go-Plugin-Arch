package main

import (
	"context"
	"fmt"
	"time"
)

func Load(ctx context.Context) {
	fmt.Println("plugin3 is loaded.")
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the main loop
			fmt.Println("plugin 3 main loop exiting due to cancellation signal.")
			return
		default:
			fmt.Println("hello world from plugin 3")
			time.Sleep(time.Second)
		}
	}

}
