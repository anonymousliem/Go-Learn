package Goroutine_Learn

import (
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0
	for i:=1;i<=1000;i++{
		go func() {
			for j:=1;j<=100;j++{
				x = x +1
			}
 		}()
	}
	time.Sleep(5*time.Second)
	println("counter = ", x)
}

