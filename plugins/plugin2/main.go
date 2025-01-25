package main

import "fmt"

func Load() {
	fmt.Println("plugin2 is loaded.")
	for range 10 {
		fmt.Println("hello world from plugin 2")
	}
}
