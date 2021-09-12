package cashier

import (
	"github.com/beecorrea/midas/src/entities/external"
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/beecorrea/midas/src/services/public"
	"github.com/google/uuid"
)


type Cashier interface {
	TransferMoney(owner *external.AccountOwner, to_acc uuid.UUID, amount int) int
}

type CashierProc struct {
	 pb *public.PublicServices
}

func NewCashier(pb *public.PublicServices) *CashierProc {
	return &CashierProc{pb: pb}
}

func (c *CashierProc) TransferMoney(owner *external.AccountOwner, to_acc uuid.UUID, amount int) int {
	
	to := c.pb.Cashier.FetchAccount(to_acc)
	// TODO refator to NewTransfer in PrivateCashierService
	tf := &private.Transfer{
		From: owner.Account,
		To: to,
		Amount: amount,
	}
	
	c.pb.Cashier.TransferMoney(tf)
	return tf.Amount
}

