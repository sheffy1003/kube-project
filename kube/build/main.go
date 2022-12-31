package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello suarez world!")

		time.Sleep(time.Second * 1)
	}
}
