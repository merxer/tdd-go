package wallet

type Wallet struct {
	cash int
}

func (w *Wallet) Deposit(cash_in int) {
	w.cash = cash_in
}

func (w *Wallet) Balance() (balance int) {
	return w.cash
}
