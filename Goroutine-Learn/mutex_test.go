package Goroutine_Learn

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexRaceCondition(t *testing.T) {
	var mutex sync.Mutex
	x := 0
	for i:=1;i<=1000;i++{
		go func() {
			for j:=1;j<=100;j++{
				mutex.Lock()
				x = x +1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5*time.Second)
	println("counter = ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int{
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestMutexReadWrite(t *testing.T) {
	account := BankAccount{}

	for i:=0;i<100;i++{
		go func() {
			for j:=0;j<100;j++{
				account.AddBalance(i)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5*time.Second)
	fmt.Println("total balance", account.GetBalance())
}

type UserBallance struct {
	sync.Mutex
	Name string
	Ballance int
}

func (user *UserBallance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBallance) Unlock()  {
	user.Mutex.Unlock()
}

func (user *UserBallance) Change(amount int)  {
	user.Ballance = user.Ballance + amount
}

func Transfer(user1 *UserBallance, user2 *UserBallance, amount int){
	user1.Lock()
	user1.Change(-amount)
	fmt.Println("Lock User 1", user1.Name)

	time.Sleep(1*time.Second)

	user2.Lock()
	user2.Change(amount)

	fmt.Println("Lock User 2", user2.Name)
	time.Sleep(1*time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T)  {
	user1 := UserBallance{
		Name: "Jokowi",
		Ballance: 1000,
	}

	user2 := UserBallance{
		Name: "Prabowo",
		Ballance: 1000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(3*time.Second)
	fmt.Println("User 1 ", user1.Name, ", Balance ", user1.Ballance)

	fmt.Println("User 2 ", user2.Name, ", Balance ", user2.Ballance)
}