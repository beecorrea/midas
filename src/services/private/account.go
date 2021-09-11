package private

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

type PrivaceAccountService struct {}

func NewPrivateAccountSvc() *PrivaceAccountService {
	return &PrivaceAccountService{}
}

func (pas *PrivaceAccountService) NewAccount(owner uuid.UUID) *private.Account {
	act := &private.Account{
		Id: uuid.New(),
		Owner: owner,
		Funds: 0,
	}

	return act
}

func (pas *PrivaceAccountService) DeleteAccount(acc *private.Account) int{
	funds := pas.DrawMoney(acc, acc.Funds)
	acc = nil
	return funds
}

func (pas PrivaceAccountService) DrawMoney(acc *private.Account, amount int) int {
	acc.Funds -= amount
	return amount
}

func (pas PrivaceAccountService) DepositMoney(acc *private.Account, amount int) {
	acc.Funds += amount
}

func (pas PrivaceAccountService) GetFunds(acc *private.Account) int{
	return acc.Funds
}




