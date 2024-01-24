package handlers

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// Set up the upgrader with necessary parameters
	// For instance, you can check the origin of the request like this:
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func HandleHotReloadWS(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// Watch for changes in the templates folder
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					ws.WriteMessage(websocket.TextMessage, []byte("reload"))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("../templates")
	if err != nil {
		log.Fatal(err)
	}

	// Keep the connection open
	for {
		if _, _, err := ws.NextReader(); err != nil {
			ws.Close()
			break
		}
	}
}