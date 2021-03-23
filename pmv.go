package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func move(wg *sync.WaitGroup, src string, dest string, overwrite bool) {
	if _, err := os.Stat(src); err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Can't stat source: ", src))
		return
	}
	if _, err := os.Stat(dest); err == nil && !overwrite {
		os.Stderr.WriteString(fmt.Sprintln("Destination exists: ", dest))
		return
	}
	err := os.Rename(src, dest)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Can't rename:", err))
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	dest, files := os.Args[len(os.Args)-1], os.Args[1:len(os.Args)-1]

	//TODO: flags

	if _, err := os.Stat(dest); err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Destination folder doesn't exist: ", err))
		os.Exit(1)
	}

	for _, file := range files {
		wg.Add(1)
		file := file
		go move(&wg, file, filepath.Join(dest, filepath.Base(file)), false)
	}
	wg.Wait()
}
