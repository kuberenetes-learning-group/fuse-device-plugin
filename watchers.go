package main

import (
	"log"
	"github.com/fsnotify/fsnotify"
	"os"
	"os/signal"
)

func newFSWatcher(files ...string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Errorf("new watcher meet error:%s\n", err.Error())
		return nil, err
	}

	for _, f := range files {
		err = watcher.Add(f)
		if err != nil {
			log.Errorf("add watcher meet error:%s\n", err.Error())
			watcher.Close()
			return nil, err
		}
	}

	return watcher, nil
}

func newOSWatcher(sigs ...os.Signal) chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)

	return sigChan
}
