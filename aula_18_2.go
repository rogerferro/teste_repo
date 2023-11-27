package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Número CPU: ", runtime.NumCPU())
	fmt.Println("Número de Go Rotines: ", runtime.NumGoroutine())

	wg.Add(1)

	go func1()
	go func2()

	fmt.Println("Número de Go Rotines: ", runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(100)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(50)
	}
	wg.Done()
}
