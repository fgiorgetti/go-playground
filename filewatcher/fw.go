package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fgiorgetti/go-playground/filewatcher/pkg/watcher"
)

type MyHandler struct {
}

func (m *MyHandler) OnCreate(name string) {
	log.Printf("File has been created: %s", name)
}

func (m *MyHandler) OnUpdate(name string) {
	log.Printf("File has been updated: %s", name)
}

func (m *MyHandler) OnRemove(name string) {
	log.Printf("File has been removed: %s", name)
}

func main() {
	stopCh := make(chan bool)
	if len(os.Args) != 2 {
		log.Fatalf("Use: %s file_or_directory", os.Args[0])
	}
	fileOrDir := os.Args[1]
	err := watcher.NewWatcher(fileOrDir, stopCh, &MyHandler{})
	if err != nil {
		log.Fatalf("Error creating watcher: %v", err)
	}

	// var done string
	fmt.Println("Press ENTER when done")
	_, _ = fmt.Scanln()
	close(stopCh)
	time.Sleep(time.Second)
}
