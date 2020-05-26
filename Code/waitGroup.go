package main

import (
	"fmt"
	"sync"
	"time"
)

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("here")
	}
}

func printStuff(wg *sync.WaitGroup) {
	// defer (*wg).Done()
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

func mainwg() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printStuff(&wg)
	wg.Wait()
}
