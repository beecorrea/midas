package processes

import (
	acc "github.com/beecorrea/midas/src/processes/account_manager"
	"github.com/beecorrea/midas/src/processes/cashier"
	"github.com/beecorrea/midas/src/services/public"
)

type Cashier cashier.Cashier
func GetCashier(pb *public.PublicServices) Cashier {
	return cashier.NewCashier(pb)
}

type AccountManager acc.AccountManager
func GetAccManager(pb *public.PublicServices) AccountManager {
	return acc.NewAccountManager(pb)
}