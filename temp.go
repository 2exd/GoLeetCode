package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 6; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
