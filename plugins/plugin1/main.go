package main

import "fmt"

func load() {
	fmt.Println("plugin1 is loaded.")
	for range 10 {
		fmt.Println("hello world from plugin 1")
	}
}
