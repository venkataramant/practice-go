package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("first func %d\n", i)
			time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			fmt.Printf("second func %d\n", i)
			time.Sleep(time.Duration(rand.Intn(200) * int(time.Millisecond)))
		}
	}()

	fmt.Println("Main Goin to wait")
	wg.Wait()
	fmt.Println("End of Main")
}
