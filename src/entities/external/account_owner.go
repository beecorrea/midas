package external

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

type AccountOwner struct {
	Id uuid.UUID
	Account *private.Account
}

