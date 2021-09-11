package public

import (
	"github.com/beecorrea/midas/src/entities/private"
	pvs "github.com/beecorrea/midas/src/services/private"

	"github.com/google/uuid"
)


type PublicAccManagerSvc struct {
	pv pvs.PrivateServices
}

func NewPublicAccManagerSvc(prv pvs.PrivateServices) *PublicAccManagerSvc{
	return &PublicAccManagerSvc{
		pv: prv,
	}
}

func (pams *PublicAccManagerSvc) OpenAccount(owner uuid.UUID) *private.Account {
	act := &private.Account{
		Id: uuid.New(),
		Owner: owner,
		Funds: 0,
	}

	return act
}

