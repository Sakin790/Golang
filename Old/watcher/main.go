package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// 1. Initialize a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Failed to create watcher:", err)
	}
	defer watcher.Close()

	// 2. Start a goroutine to receive and process events in the background
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Triggers on any file operation
				log.Printf("Event: %s | File: %s", event.Op, event.Name)

				// Track specific file operations
				if event.Has(fsnotify.Write) {
					log.Println("--> File was modified/written!")
				}
				if event.Has(fsnotify.Create) {
					log.Println("--> New file was created!")
				}
				if event.Has(fsnotify.Remove) {
					log.Println("--> File was deleted!")
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error received:", err)
			}
		}
	}()

	// 3. Add the target file or directory path to the watcher
	// Monitoring a log file in the current directory as an example
	err = watcher.Add("./logs")
	if err != nil {
		log.Fatal("Failed to add file to watch list:", err)
	}

	log.Println("Watching app.log continuously... (Press Ctrl+C to stop)")
	<-done // Keeps the script running indefinitely
}
