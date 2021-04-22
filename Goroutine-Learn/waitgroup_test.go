package Goroutine_Learn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Waitasync(group *sync.WaitGroup){
	defer  group.Done()

	group.Add(1)

	fmt.Println("Halo")
	time.Sleep(1*time.Second)
}

func TestWaitAsync(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i<100;i++{
		go Waitasync(group)
	}

	group.Wait()
	fmt.Println("selesai")
}