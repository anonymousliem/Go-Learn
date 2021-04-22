package Goroutine_Learn

import (
	"fmt"
	"testing"
	"time"
)

func HellWorld()  {
	fmt.Println("Hello World")
}

func TestCreateGoRoutines(t *testing.T) {
	go HellWorld()
	fmt.Println("okay")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i:=0; i < 100000;i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}