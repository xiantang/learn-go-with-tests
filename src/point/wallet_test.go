package point

import (
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	if got == nil {
		t.Fatal("wanted an error but didnt get one")
	}
	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}
