package main

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

func handleRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func mainerr() {
	defer handleRecover()

	if err := openFile(); err != nil {
		panic(fmt.Sprintln("Somethign happened", err))
	}

}

func openFile() error {
	f, err := os.Open("anyfile.txt")
	if err != nil {
		return err
	}
	defer f.Close() // after successful open
	return nil
}
