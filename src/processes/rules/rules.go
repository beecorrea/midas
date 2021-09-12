package rules

import (
	ent "github.com/beecorrea/midas/src/entities/private"
	"github.com/beecorrea/midas/src/services/private"
	"github.com/google/uuid"
)

func DoesAccountExist(account *ent.Account) bool{
	if account == nil {
		return false
	}

	for _, acc := range private.GetAllAccounts() {
		if acc.Id == account.Id {
			return true
		}
	}

	return false
}

func IsOwnerOfAccount(acc *ent.Account, ownerId uuid.UUID) bool {
	return acc.Owner == ownerId
}

func CanDraw(funds int, drawRequest int) bool {
	return drawRequest <= funds
}