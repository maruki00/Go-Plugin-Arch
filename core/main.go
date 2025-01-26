package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"plugin"
	"time"
)

func main() {
	loadedPlugins := make(map[string]context.CancelFunc, 0)
	defer func() {
		// Recover from panic
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	for {
		trashPlug, err := os.ReadDir("./trash")
		if err == nil {
			for _, tsh := range trashPlug {
				plug, ok := loadedPlugins[tsh.Name()]
				if !ok {
					continue
				}
				plug()
				os.Remove(path.Join("./trash", tsh.Name()))
				delete(loadedPlugins, tsh.Name())
			}
		}

		pluginsDir, err := os.ReadDir("./plugins")
		if err != nil {
			fmt.Println("error read plugins folder : ", err.Error())
			continue
		}
		for _, Plugfile := range pluginsDir {
			if _, ok := loadedPlugins[Plugfile.Name()]; ok {
				continue
			}

			plug, err := plugin.Open(path.Join("plugins", Plugfile.Name()))
			if err != nil {
				fmt.Println("plugin Open : ", err.Error())
				continue
			}
			psym, err := plug.Lookup("Load")
			if err != nil {
				fmt.Println("Lookup : could not load the plugin : ", Plugfile.Name())
				continue
			}

			LoadFunc, ok := psym.(func(ctx context.Context))
			if !ok {
				fmt.Println("Invalid function signature")
				return
			}

			ctx, cancel := context.WithCancel(context.Background())
			go LoadFunc(ctx)
			loadedPlugins[Plugfile.Name()] = cancel
		}
		time.Sleep(time.Second)
	}
}
