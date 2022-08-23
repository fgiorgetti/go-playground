package watcher

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

type FSChangeHandler interface {
	OnCreate(string)
	OnUpdate(string)
	OnRemove(string)
}

func watchCreated(watcher *fsnotify.Watcher, name string) {
	log.Printf("-> waiting for %s to exist", name)
	go func() {
		ticker := time.Tick(time.Second)
		for {
			select {
			case <-ticker:
				if err := watcher.Add(name); err == nil {
					log.Printf("-> now it exists: %s", name)
					return
				}
			}
		}
	}()
	return
}

func NewWatcher(name string, stopCh chan bool, handler FSChangeHandler) error {
	log.Printf("Creating watcher for: %s", name)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	if _, err := os.Stat(name); err != nil && os.IsNotExist(err) {
		watchCreated(watcher, name)
	} else {
		watcher.Add(name)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				// useful for new files when watching directories
				case event.Op&fsnotify.Create == fsnotify.Create:
					handler.OnCreate(event.Name)
				case event.Op&fsnotify.Write == fsnotify.Write:
					handler.OnUpdate(event.Name)
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					handler.OnRemove(event.Name)
					// object being watched removed, watch for it to show up again
					if event.Name == name {
						watcher.Remove(name)
						watchCreated(watcher, name)
					}
				}
			case <-stopCh:
				log.Printf("Done watching: %s", name)
				watcher.Close()
				return
			}
		}
	}()

	return nil
}
