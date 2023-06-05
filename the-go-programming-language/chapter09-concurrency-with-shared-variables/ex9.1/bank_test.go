package main

import "testing"

type Action struct {
	Deposit    int
	Withdrawal int
}

func TestBan(t *testing.T) {
	tcs := []struct {
		action Action
		want   bool
	}{
		{Action{100, 100}, true},
		{Action{200, 100}, true},
		{Action{50, 100}, false},
	}
	for _, tc := range tcs {
		account := NewAccount()
		account.Deposit(tc.action.Deposit)
		ok := account.Withdraw(tc.action.Withdrawal)
		if ok != tc.want {
			t.Errorf("Deposit %d Withdraw %d got %t, want %t", tc.action.Deposit, tc.action.Withdrawal, ok, tc.want)
		}
	}
}
