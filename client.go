package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

type APIClient struct {
	address string
	service Service
	client  http.Client
}

func NewAPIClient(address string, service Service) *APIClient {
	return &APIClient{
		address: address,
		service: service,
		client:  http.Client{Timeout: time.Duration(15) * time.Second},
	}
}

func (client *APIClient) Run() {
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for
	// directories
	if err := filepath.Walk("/home/luqus/Documents/test", watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				// fmt.Println(event.Name)

				// client.service.UploadFile()
				client.uploadFile(event.Name)

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func (client *APIClient) uploadFile(path string) {
	client.service.UploadFile(path)

}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}
