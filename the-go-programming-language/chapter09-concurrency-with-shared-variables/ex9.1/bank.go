// Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
// program. The result should indicate whether the transaction succeeded or failed
// due to insufficiet funds. The message sent to the monitor goroutine must contain
// both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result
// back to Withdraw
// copied torbiak/gopl/ex9.1/bank.go
package main

type Withdrawal struct {
	amount  int
	success chan bool
}

type Account struct {
	deposits chan int
	balances chan int
	withdraw chan Withdrawal
}

func NewAccount() *Account {
	deposits := make(chan int)        // send amount to deposit
	balances := make(chan int)        // receive balance
	withdraw := make(chan Withdrawal) // deduct amount
	account := &Account{deposits, balances, withdraw}
	go account.teller()
	return account
}

func (a *Account) Deposit(amount int) { a.deposits <- amount }
func (a *Account) Balance() int       { return <-a.balances }
func (a *Account) Withdraw(amount int) bool {
	ch := make(chan bool)
	a.withdraw <- Withdrawal{amount, ch}
	return <-ch
}

func (a *Account) teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-a.deposits:
			balance += amount
		case a.balances <- balance:
		case withdrawal := <-a.withdraw:
			if withdrawal.amount <= balance {
				balance -= withdrawal.amount
				withdrawal.success <- true
			} else {
				withdrawal.success <- false
			}
		}
	}
}
