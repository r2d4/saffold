package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello new deploy!")
		time.Sleep(time.Second * 1)
	}
}
