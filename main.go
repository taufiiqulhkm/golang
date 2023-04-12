package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func processData1(data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println("Processing data1:", data)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func processData2(data interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println("Processing data2:", data)
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	data1 := "bisa1"
	data2 := "bisa2"

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(2)
		go processData1(data1, &wg)
		go processData2(data2, &wg)
	}
	wg.Wait()

	fmt.Println("")

	for i := 0; i < 4; i++ {
		var wg2 sync.WaitGroup
		wg2.Add(2)
		go func() {
			processData1(data1, &wg2)
		}()
		go func() {
			processData2(data2, &wg2)
		}()
		wg2.Wait()
	}

	time.Sleep(5 * time.Second)
}