package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync/atomic"
	"time"

	"github.com/fgiorgetti/go-playground/filewatcher/pkg/watcher"
)

type MyHandler struct {
	created uint32
	updated uint32
	deleted uint32
}

func (m *MyHandler) OnCreate(name string) {
	log.Printf("File has been created: %s", name)
	atomic.AddUint32(&m.created, 1)
}

func (m *MyHandler) OnUpdate(name string) {
	log.Printf("File has been updated: %s", name)
	atomic.AddUint32(&m.updated, 1)
}

func (m *MyHandler) OnRemove(name string) {
	log.Printf("File has been removed: %s", name)
	atomic.AddUint32(&m.deleted, 1)
}

func main() {
	stopCh := make(chan struct{})
	//if len(os.Args) != 2 {
	//	log.Fatalf("Use: %s file_or_directory", os.Args[0])
	//}
	//fileOrDir := os.Args[1]

	w, err := watcher.NewWatcher()
	if err != nil {
		log.Fatalf("error creating watcher: %s", err)
	}
	h := &MyHandler{}
	w.Add(os.TempDir(), h, regexp.MustCompile(`/fernando\.[0-9]0*$`))
	w.Start(stopCh)

	// display totals
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				log.Printf("Created: %d - Updated: %d - Deleted: %d", h.created, h.updated, h.deleted)
			case <-stopCh:
			}
		}
	}()
	// var done string
	fmt.Println("Press ENTER when done")
	_, _ = fmt.Scanln()
	close(stopCh)
	time.Sleep(time.Second)
}
