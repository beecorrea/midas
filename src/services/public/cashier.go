package public

import (
	"github.com/beecorrea/midas/src/entities/private"
	pvs "github.com/beecorrea/midas/src/services/private"
)

type PublicCashierService struct {
	pv pvs.PrivateServices
}

func NewPublicCashierService(prv pvs.PrivateServices) *PublicCashierService{
	return &PublicCashierService{
		pv: prv,
	}
}

func (pcs *PublicCashierService) CheckAccountFunds(acc *private.Account) int {
	return pcs.pv.Account.GetFunds(acc)
}

func (pcs *PublicCashierService) Draw(acc *private.Account, amount int) int {
	return pcs.pv.Account.DrawMoney(acc, amount)
}

func (pcs *PublicCashierService) Deposit(acc *private.Account, amount int) {
	pcs.pv.Account.DepositMoney(acc, amount)
}

func (pcs *PublicCashierService) TransferMoney(from *private.Account, to *private.Account, amount int) {
	amtToTransfer := pcs.pv.Account.DrawMoney(from, amount)
	pcs.pv.Account.DepositMoney(to, amtToTransfer)
}