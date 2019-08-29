package point

import (
	"fmt"
	"github.com/pkg/errors"
)

//Go 允许从现有的类型创建新的类型。
type Bitcoin int

// var 关键字允许我们定义包的全局值。
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// 当调用时 w是我们来自调用方法的副本
func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
