package public

import (
	"github.com/beecorrea/midas/src/entities/private"
	pvs "github.com/beecorrea/midas/src/services/private"

	"github.com/google/uuid"
)


type PublicAccManagerSvc struct {
	pv *pvs.PrivateServices
}

func NewPublicAccManagerSvc(prv *pvs.PrivateServices) *PublicAccManagerSvc{
	return &PublicAccManagerSvc{
		pv: prv,
	}
}

func (pams *PublicAccManagerSvc) CreateAccount(owner uuid.UUID) *private.Account {
	act := pams.pv.Account.NewAccount(owner)
	return act
}

func (pcs *PublicAccManagerSvc) FetchAccount(id uuid.UUID) *private.Account {
	return pcs.pv.Account.GetAccountByID(id)
}