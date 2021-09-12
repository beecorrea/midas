package private

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

var GlobalAccounts []*private.Account = make([]*private.Account, 0)

type PrivateAccountService struct {}

func NewPrivateAccountSvc() *PrivateAccountService {
	return &PrivateAccountService{}
}

func (pas *PrivateAccountService) NewAccount(owner uuid.UUID) *private.Account {
	act := &private.Account{
		Id: uuid.New(),
		Owner: owner,
		Funds: 0,
	}
	GlobalAccounts = append(GlobalAccounts, act)
	return act
}

func (pas *PrivateAccountService) DeleteAccount(acc *private.Account) int{
	funds := pas.DrawMoney(acc, acc.Funds)
	acc = nil
	return funds
}

func (pas PrivateAccountService) DrawMoney(acc *private.Account, amount int) int {
	acc.Funds -= amount
	return amount
}

func (pas PrivateAccountService) DepositMoney(acc *private.Account, amount int) {
	acc.Funds += amount
}

func (pas PrivateAccountService) GetFunds(acc *private.Account) int{
	return acc.Funds
}

func (pas PrivateAccountService) GetAccountByID(id uuid.UUID) *private.Account {
	for _, acc := range GlobalAccounts {
		if(acc.Id == id) {
			return acc
		}
	}
	return pas.NewAccount(id)
}

func GetAllAccounts() []*private.Account{
	return GlobalAccounts
}

