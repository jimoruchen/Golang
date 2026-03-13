package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println(0)
			ch <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-ch
			fmt.Println(1)
			ch <- struct{}{}
		}
	}()
	ch <- struct{}{}
	wg.Wait()
}
