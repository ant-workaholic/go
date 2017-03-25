package funding

type Fund struct {
    balance int
}

// A regular function returning a point to fund
func NewFund(initialBalance int) *Fund {
    return &Fund{
        balance: initialBalance,
    }
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
    return f.balance
}

func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}
