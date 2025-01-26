package main

import (
	"fmt"
	"time"
)

func Load() {
	fmt.Println("plugin3 is loaded.")
	for {
		fmt.Println("hello world from plugin 3")
		time.Sleep(time.Second)
	}
}
