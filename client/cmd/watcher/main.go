package main

import (
	"log"
	"os/user"
	"strings"

	"github.com/fsnotify/fsnotify"
)

const (
	watchDirPath = "~/Music/djay/djay Media Library.djayMediaLibrary"
	watchFile    = "NowPlaying.txt"
	backendAddr  = "http://localhost:8080"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to create wather: %v", err)
	}
	defer watcher.Close()

	usr, _ := user.Current()
	dir := strings.Replace(watchDirPath, "~", usr.HomeDir, 1)

	err = watcher.Add(dir)
	if err != nil {
		log.Fatalf("failed to watch target directory: %v", err)
	}

	log.Println("start watching...", dir)
	for {
		select {
		case event := <-watcher.Events:
			if strings.Contains(event.Name, watchFile) && event.Op == fsnotify.Create {
				log.Println("modified file:", event.Name, event.Op)
			}
			if strings.Contains(event.Name, watchFile) && event.Op == fsnotify.Remove {
				log.Println("modified file:", event.Name, event.Op)
			}
		case err := <-watcher.Errors:
			log.Printf("error: %v", err)
		}
	}
}
