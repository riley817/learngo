package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStatus(mem runtime.MemStats)  {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc : ", mem.Alloc)
	fmt.Println("mem.TotalAlloc : ", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc : ", mem.HeapAlloc)
	fmt.Println("mem.NumGC : ", mem.NumGC)
	fmt.Println("--------")
}

func main() {
	var mem runtime.MemStats
	printStatus(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStatus(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStatus(mem)

}
