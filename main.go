package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello new commit!")
		time.Sleep(time.Second * 1)
	}
}
