package main

/**
 * uses fsnotify to listen for change events; an event triggers an action as supplied on the command line,
 * e.g., watcher --command "go run file.go"
 */

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"strings"
	"time"
)

var (
	watchdir string
	action   string
	suffix   string
)

func configForGo() {
	action = "go run %s"
	suffix = ".go"
}

func main() {
	watchdir = "../util"
	configForGo()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	files := make(map[string]int64)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, ".go") {
					files[event.Name] = time.Now().Unix()
					fmt.Println("event:", event, files)
				}
			case err := <-watcher.Errors:
				fmt.Printf("error %v\n", err)
			}
		}
	}()

	fmt.Printf("watch folder %s\n", dir)
	err = watcher.Add(dir)
	if err != nil {
		panic(err)
	}

	<-done

	fmt.Println("all done...")
}
