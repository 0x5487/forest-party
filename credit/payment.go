package wallet

type Transaction struct {
	ID string
}

type WalletServicer interface {
	Withdraw(credit int64) error
	Deposit(credit int64) error
}
