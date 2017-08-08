package main

import (
	"fmt"
	"sync"
	"time"
)

var m = map[string]int{"a": 1}
var lock = sync.RWMutex{}

func main() {
	go Read()
	time.Sleep(1 * time.Second)
	fmt.Println(m)
	go Write()
	fmt.Println(m)
	time.Sleep(1 * time.Minute)
	fmt.Println(m)
}

// Read from a map
func Read() {
	for {
		read()
		fmt.Println(m)
	}
}

// Write to a map
func Write() {
	for {
		write()
		fmt.Println(m)
	}
}

func read() {
	lock.RLock()
	defer lock.RUnlock()
	_ = m["a"]
}

func write() {
	lock.Lock()
	defer lock.Unlock()
	m["b"] = 2
}
