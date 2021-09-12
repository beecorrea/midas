package cashier

import (
	"github.com/beecorrea/midas/src/entities/external"
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/beecorrea/midas/src/processes/rules"
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
	ownerHasAccount := rules.DoesAccountExist(owner.Account)
	
	if !ownerHasAccount {
		return -1
	}
	
	if ownerThemself := rules.IsOwnerOfAccount(owner.Account, owner.Id); !ownerThemself {
		return -2
	}

	to := c.pb.Cashier.FetchAccount(to_acc)
		
	if destHasAccount := rules.DoesAccountExist(to); !destHasAccount {
		return -3
	}

	if canDraw := rules.CanDraw(owner.Account.Funds, amount); !canDraw {
		return -4
	}

	// TODO refator to NewTransfer in PrivateCashierService
	tf := &private.Transfer{
		From: owner.Account,
		To: to,
		Amount: amount,
	}
	
	c.pb.Cashier.TransferMoney(tf)
	return tf.Amount
}

