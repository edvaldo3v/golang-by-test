package pointer_error

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want int) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	}

	t.Run("Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := 10
		assertBalance(t, wallet, want)
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)
		if err != nil {
			t.Errorf("Error %v", err)
		}
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw nsuficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, 20)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
