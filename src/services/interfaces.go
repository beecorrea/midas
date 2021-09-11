package services

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

type Account interface {
	NewAccount(owner uuid.UUID) *private.Account
	DeleteAccount(acc *private.Account) int
	DrawMoney(acc *private.Account, amount int) int
	DepositMoney(acc *private.Account, amount int) 
	GetFunds(acc *private.Account) int
}

type Transfer interface {
	NewTransfer(from *private.Account, to *private.Account, amount int) *private.Transfer
}

type AccountManager interface {
	OpenAccount(owner uuid.UUID) *private.Account
	CloseAccount(acc *private.Account)
}

type Cashier interface {
	CheckAccountFunds(acc *private.Account) int
	Draw(acc *private.Account, amount int) int
	Deposit(acc *private.Account, amount int)
	TransferMoney(from uuid.UUID, to uuid.UUID, amount int)
}