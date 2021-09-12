package acc

import (
	"github.com/beecorrea/midas/src/entities/external"
	"github.com/beecorrea/midas/src/processes/rules"
	"github.com/beecorrea/midas/src/services/public"
	"github.com/google/uuid"
)

type AccountManager interface {
	OpenAccount(owner *external.AccountOwner) uuid.UUID
}

type AccountManagerProc struct {
	 pb *public.PublicServices
}

func NewAccountManager(pb *public.PublicServices) *AccountManagerProc {
	return &AccountManagerProc{pb: pb}
}

func (amp *AccountManagerProc) OpenAccount(owner *external.AccountOwner) uuid.UUID {
	if accountExists := rules.DoesAccountExist(owner.Account); accountExists == true {
		return owner.Account.Id
	}
	
	owner.Account = amp.pb.AccountManager.CreateAccount(owner.Id)
	return owner.Account.Id
}