package public

import (
	"github.com/beecorrea/midas/src/entities/private"
	"github.com/google/uuid"
)

type AccountOwner struct {
	Id uuid.UUID
	Accounts []private.Account
}

