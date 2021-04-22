package Goroutine_Learn

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			time.Sleep(3*time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total cpu ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)

	group.Wait()
}

func TestThreadChangeNumber(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)

		go func() {
			time.Sleep(3*time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total cpu ", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine", totalGoroutine)

	group.Wait()
}