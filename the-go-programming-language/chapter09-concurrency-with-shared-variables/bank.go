package main

import "sync"

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }

// gopl.io/ch9/bank1
// share memory by communicating
// provides a concurrency-safe bank with one account.
var deposits1 = make(chan int) // send amount to deposit
var balances1 = make(chan int) // receive balance

func Deposit1(amount int) { deposits1 <- amount }
func Balance1() int       { return <-balances1 }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits1:
			balance += amount
		case balances1 <- balance:
		}
	}
}

func init() {
	go teller()
}

// gopl.io/ch9/bank2
var (
	sema     = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance2 int
)

func Deposit2(amount int) {
	sema <- struct{}{} // acquire toke
	balance = balance + amount
	<-sema // release token
}

func Balance2() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema //release token
	return b
}

// gopl.io/ch9/bank3
var (
	mu       sync.Mutex // guards balance
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

// NOTE: not atomic!
// func Withdraw(amount int) bool {
// 	Deposit(-amount)
// 	if Balance() < 0 {
// 		Deposit(amount)
// 		return false // insufficient funds
// 	}
// 	return true
// }

// common solution:
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if Balance() < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

func deposit(amount int) {
	balance = balance + amount
}

var mu4 sync.RWMutex
var balance4 int

func Balance4() int {
	mu4.RLock() // readers lock
	defer mu4.RUnlock()
	return balance
}
