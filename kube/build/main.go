package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello New world!")

		time.Sleep(time.Second * 1)
	}
}
