package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Error("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func ExampleBitcoin_String() {
	str := Bitcoin.String(Bitcoin(10))
	fmt.Println(str)
	// Output: 10 BTC
}

func ExampleWallet_Deposit() {
	wallet := Wallet{}
	wallet.Deposit(10)
	fmt.Println(wallet.balance)
	// Output: 10 BTC
}

func ExampleWallet_Balance() {
	wallet := Wallet{balance: 20}
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 20 BTC
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{balance: 30}
	wallet.Withdraw(20)
	fmt.Println(wallet.balance)
	// Output: 10 BTC
}

func ExampleWallet_Withdraw_second() {
	wallet := Wallet{balance: 20}
	err := wallet.Withdraw(100)
	fmt.Println(err)
	// Output: cannot withdraw, insufficient funds
}
