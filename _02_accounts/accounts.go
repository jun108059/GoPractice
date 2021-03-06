package _02_accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string // 계좌 주인 (private)
	balance int    // 잔액 (private)
}

var errNoMoney = errors.New("can't withdraw you are poor")

// NewAccount Account 생성 (생성자로 동작)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit 계좌에 돈 입금하는 method
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// GetBalance 계좌 잔액 불러오기
func (a Account) GetBalance() int {
	return a.balance
}

// Withdraw 계좌에서 돈을 출금
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner 계좌 주인 이름 변경
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// GetOwner 계좌 주인 이름 가져오기
func (a Account) GetOwner() string {
	return a.owner
}

// String python __str__ 매직메소드처럼 호출 됨
func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account.\n Has: ", a.GetBalance())
}
