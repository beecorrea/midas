package interfaces

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

type Account interface {
	NewAccount(owner uuid.UUID) *private.Account
	DeleteAccount(acc *private.Account) int
	GetAccountByID(id uuid.UUID) *private.Account
	DrawMoney(acc *private.Account, amount int) int
	DepositMoney(acc *private.Account, amount int) 
	GetFunds(acc *private.Account) int
}

type Transfer interface {
	NewTransfer(from *private.Account, to *private.Account, amount int) *private.Transfer
}

type AccountManager interface {
	CreateAccount(owner uuid.UUID) *private.Account
}

type Cashier interface {
	CheckAccountFunds(acc *private.Account) int
	Draw(acc *private.Account, amount int) int
	Deposit(acc *private.Account, amount int)
	TransferMoney(tf *private.Transfer)
	FetchAccount(id uuid.UUID) *private.Account
}