package public

import (
	"github.com/beecorrea/midas/src/entities/private"

	"github.com/google/uuid"
)

func OpenAccount(owner uuid.UUID) *private.Account {
	act := &private.Account{
		Id: uuid.New(),
		Owner: owner,
		Funds: 0,
	}

	return act
}

