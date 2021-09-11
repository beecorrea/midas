package public

import (
	"github.com/beecorrea/midas/src/entities/private"
	pvs "github.com/beecorrea/midas/src/services/private"
	"github.com/google/uuid"
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

func (pcs *PublicCashierService) TransferMoney(tf *private.Transfer) {
	amtToTransfer := pcs.pv.Account.DrawMoney(tf.From, tf.Amount)
	pcs.pv.Account.DepositMoney(tf.To, amtToTransfer)
}

func (pcs *PublicCashierService) FetchAccount(id uuid.UUID) *private.Account {
	return pcs.pv.Account.GetAccountByID(id)
}